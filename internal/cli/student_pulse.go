package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/spf13/cobra"
)

type studentEnrollment struct {
	UserID int `json:"user_id"`
	User   struct {
		Name string `json:"name"`
	} `json:"user"`
	Grades struct {
		CurrentScore *float64 `json:"current_score"`
	} `json:"grades"`
}

type studentSubmission struct {
	UserID        int     `json:"user_id"`
	SubmittedAt   *string `json:"submitted_at"`
	WorkflowState string  `json:"workflow_state"`
}

func newStudentPulseCmd(flags *rootFlags) *cobra.Command {
	var courseID string
	var atRiskDays int
	cmd := &cobra.Command{
		Use:   "student-pulse",
		Short: "Grade distribution and at-risk students for a course",
		Long:  "Shows grade distribution (A/B/C/D/F) and flags students who haven't submitted recently.",
		RunE: func(cmd *cobra.Command, args []string) error {
			cc, err := newCanvasClientFromConfig(flags.configPath)
			if err != nil {
				return err
			}
			return runStudentPulse(cc, courseID, atRiskDays, cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVarP(&courseID, "course", "c", "", "Course ID (required)")
	cmd.Flags().IntVar(&atRiskDays, "at-risk-days", 14, "Days without submission to flag as at-risk")
	_ = cmd.MarkFlagRequired("course")
	return cmd
}

func runStudentPulse(cc *canvasClient, courseID string, atRiskDays int, w io.Writer) error {
	// Fetch enrollments
	enrollRaw, err := cc.getAll(fmt.Sprintf("/courses/%s/enrollments", courseID), nil)
	if err != nil {
		return fmt.Errorf("fetching enrollments: %w", err)
	}
	if len(enrollRaw) == 0 {
		fmt.Fprintln(w, "No students enrolled in this course.")
		return nil
	}

	var enrollments []studentEnrollment
	for _, raw := range enrollRaw {
		var e studentEnrollment
		if err := json.Unmarshal(raw, &e); err == nil && e.UserID != 0 {
			enrollments = append(enrollments, e)
		}
	}
	if len(enrollments) == 0 {
		fmt.Fprintln(w, "No students enrolled in this course.")
		return nil
	}

	// Fetch submissions
	subRaw, err := cc.getAll(fmt.Sprintf("/courses/%s/students/submissions", courseID), nil)
	if err != nil {
		return fmt.Errorf("fetching submissions: %w", err)
	}
	// Map user_id -> latest submitted_at
	lastSubmit := map[int]time.Time{}
	for _, raw := range subRaw {
		var s studentSubmission
		if err := json.Unmarshal(raw, &s); err != nil {
			continue
		}
		if s.SubmittedAt == nil {
			continue
		}
		t, err := time.Parse(time.RFC3339, *s.SubmittedAt)
		if err != nil {
			continue
		}
		if existing, ok := lastSubmit[s.UserID]; !ok || t.After(existing) {
			lastSubmit[s.UserID] = t
		}
	}

	// Grade distribution
	buckets := map[string]int{"A": 0, "B": 0, "C": 0, "D": 0, "F": 0, "Ungraded": 0}
	bucketOrder := []string{"A", "B", "C", "D", "F", "Ungraded"}
	for _, e := range enrollments {
		bucket := gradeBucket(e.Grades.CurrentScore)
		buckets[bucket]++
	}
	fmt.Fprintf(w, "STUDENT PULSE — Course %s\n\n", courseID)
	fmt.Fprintln(w, "Grade Distribution:")
	for _, b := range bucketOrder {
		fmt.Fprintf(w, "  %-8s %d students\n", b+":", buckets[b])
	}

	// At-risk students
	cutoff := time.Now().AddDate(0, 0, -atRiskDays)
	fmt.Fprintf(w, "\nAt-Risk Students (no submission in %d days):\n", atRiskDays)
	atRiskCount := 0
	for _, e := range enrollments {
		last, submitted := lastSubmit[e.UserID]
		if !submitted || last.Before(cutoff) {
			lastStr := "never submitted"
			if submitted {
				lastStr = "last: " + last.Format("2006-01-02")
			}
			fmt.Fprintf(w, "  %s (ID: %d) — %s\n", e.User.Name, e.UserID, lastStr)
			atRiskCount++
		}
	}
	if atRiskCount == 0 {
		fmt.Fprintln(w, "  None — all students submitted recently.")
	}
	return nil
}

func gradeBucket(score *float64) string {
	if score == nil {
		return "Ungraded"
	}
	switch {
	case *score >= 90:
		return "A"
	case *score >= 80:
		return "B"
	case *score >= 70:
		return "C"
	case *score >= 60:
		return "D"
	default:
		return "F"
	}
}
