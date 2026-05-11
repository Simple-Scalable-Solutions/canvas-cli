package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/spf13/cobra"
)

type auditFinding struct {
	Kind    string
	Message string
}

type courseAssignment struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	DueAt     *string `json:"due_at"`
	HasRubric bool    `json:"has_rubric"`
}

type courseModule struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Items []struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Published bool   `json:"published"`
	} `json:"items"`
}

type coursePage struct {
	PageID int     `json:"page_id"`
	Title  string  `json:"title"`
	Body   *string `json:"body"`
}

func newCourseAuditCmd(flags *rootFlags) *cobra.Command {
	var courseID string
	cmd := &cobra.Command{
		Use:   "course-audit",
		Short: "Scan a course for quality issues (missing due dates, rubrics, empty pages)",
		RunE: func(cmd *cobra.Command, args []string) error {
			cc, err := newCanvasClientFromConfig(flags.configPath)
			if err != nil {
				return err
			}
			return runCourseAudit(cc, courseID, cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVarP(&courseID, "course", "c", "", "Course ID (required)")
	_ = cmd.MarkFlagRequired("course")
	return cmd
}

func runCourseAudit(cc *canvasClient, courseID string, w io.Writer) error {
	var findings []auditFinding

	// Audit assignments
	assignRaw, err := cc.getAll(fmt.Sprintf("/courses/%s/assignments", courseID), nil)
	if err != nil {
		return fmt.Errorf("fetching assignments: %w", err)
	}
	for _, raw := range assignRaw {
		var a courseAssignment
		if err := json.Unmarshal(raw, &a); err != nil || a.ID == 0 {
			continue
		}
		if a.DueAt == nil {
			findings = append(findings, auditFinding{
				Kind:    "no_due_date",
				Message: fmt.Sprintf("Assignment: %q (id: %d)", a.Name, a.ID),
			})
		}
		if !a.HasRubric {
			findings = append(findings, auditFinding{
				Kind:    "no_rubric",
				Message: fmt.Sprintf("Assignment: %q (id: %d)", a.Name, a.ID),
			})
		}
	}

	// Audit modules (fetch with items included)
	params := url.Values{"include[]": {"items"}}
	moduleRaw, err := cc.getAll(fmt.Sprintf("/courses/%s/modules", courseID), params)
	if err != nil {
		return fmt.Errorf("fetching modules: %w", err)
	}
	for _, raw := range moduleRaw {
		var m courseModule
		if err := json.Unmarshal(raw, &m); err != nil || m.ID == 0 {
			continue
		}
		for _, item := range m.Items {
			if !item.Published {
				findings = append(findings, auditFinding{
					Kind:    "unpublished",
					Message: fmt.Sprintf("Module item: %q in module %q (id: %d)", item.Title, m.Name, item.ID),
				})
			}
		}
	}

	// Audit pages
	pageRaw, err := cc.getAll(fmt.Sprintf("/courses/%s/pages", courseID), nil)
	if err != nil {
		return fmt.Errorf("fetching pages: %w", err)
	}
	for _, raw := range pageRaw {
		var p coursePage
		if err := json.Unmarshal(raw, &p); err != nil || p.PageID == 0 {
			continue
		}
		if p.Body == nil || *p.Body == "" {
			findings = append(findings, auditFinding{
				Kind:    "empty_page",
				Message: fmt.Sprintf("Page: %q (id: %d)", p.Title, p.PageID),
			})
		}
	}

	// Output
	fmt.Fprintf(w, "COURSE AUDIT — Course %s\n\n", courseID)
	if len(findings) == 0 {
		fmt.Fprintln(w, "✓ No issues found.")
		return nil
	}
	fmt.Fprintf(w, "Findings (%d):\n", len(findings))
	for _, f := range findings {
		fmt.Fprintf(w, "  [%-15s] %s\n", f.Kind, f.Message)
	}
	return nil
}
