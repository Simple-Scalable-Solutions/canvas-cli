package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

type activitySummary struct {
	ID             int     `json:"id"`
	PageViews      int     `json:"page_views"`
	Participations int     `json:"participations"`
	LastActivityAt *string `json:"last_activity_at"`
}

func newActivityReportCmd(flags *rootFlags) *cobra.Command {
	var courseID string
	cmd := &cobra.Command{
		Use:   "activity-report",
		Short: "Per-student page views and participation for a course",
		RunE: func(cmd *cobra.Command, args []string) error {
			cc, err := newCanvasClientFromConfig(flags.configPath)
			if err != nil {
				return err
			}
			return runActivityReport(cc, courseID, cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVarP(&courseID, "course", "c", "", "Course ID (required)")
	_ = cmd.MarkFlagRequired("course")
	return cmd
}

func runActivityReport(cc *canvasClient, courseID string, w io.Writer) error {
	// Fetch student enrollments
	enrollRaw, err := cc.getAll(
		fmt.Sprintf("/courses/%s/enrollments", courseID),
		url.Values{"type[]": {"StudentEnrollment"}},
	)
	if err != nil {
		return fmt.Errorf("fetching enrollments: %w", err)
	}
	// Build student map: user_id -> name
	students := map[int]string{}
	var studentOrder []int
	for _, raw := range enrollRaw {
		var e studentEnrollment
		if err := json.Unmarshal(raw, &e); err != nil || e.UserID == 0 {
			continue
		}
		if _, exists := students[e.UserID]; !exists {
			students[e.UserID] = e.User.Name
			studentOrder = append(studentOrder, e.UserID)
		}
	}

	// Fetch analytics summaries
	sumRaw, err := cc.getAll(
		fmt.Sprintf("/courses/%s/analytics/student_summaries", courseID),
		nil,
	)
	if err != nil {
		return fmt.Errorf("fetching analytics: %w", err)
	}
	summaryMap := map[int]activitySummary{}
	for _, raw := range sumRaw {
		var s activitySummary
		if err := json.Unmarshal(raw, &s); err != nil || s.ID == 0 {
			continue
		}
		summaryMap[s.ID] = s
	}

	// Print report header
	fmt.Fprintf(w, "ACTIVITY REPORT — Course %s\n\n", courseID)
	fmt.Fprintf(w, "%-30s %-10s %-15s %s\n", "Student", "PageViews", "Participations", "LastActivity")
	fmt.Fprintf(w, "%-30s %-10s %-15s %s\n", strings.Repeat("-", 28), "----------", "---------------", "------------")

	// Print per-student rows
	var atRisk []string
	for _, uid := range studentOrder {
		name := students[uid]
		label := fmt.Sprintf("%s (%d)", name, uid)
		s := summaryMap[uid]
		last := "never"
		if s.LastActivityAt != nil && len(*s.LastActivityAt) >= 10 {
			last = (*s.LastActivityAt)[:10]
		}
		fmt.Fprintf(w, "%-30s %-10d %-15d %s\n", label, s.PageViews, s.Participations, last)
		if s.PageViews == 0 {
			atRiskDesc := "never active"
			if last != "never" {
				atRiskDesc = "last active: " + last
			}
			atRisk = append(atRisk, fmt.Sprintf("  %s — %s", label, atRiskDesc))
		}
	}

	// At-risk section
	fmt.Fprintf(w, "\nAt-Risk (zero page views):\n")
	if len(atRisk) == 0 {
		fmt.Fprintln(w, "  None — all students have viewed course content.")
	} else {
		for _, line := range atRisk {
			fmt.Fprintln(w, line)
		}
	}
	return nil
}
