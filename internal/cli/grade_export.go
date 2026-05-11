package cli

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type gradeSubmission struct {
	UserID       int      `json:"user_id"`
	AssignmentID int      `json:"assignment_id"`
	Score        *float64 `json:"score"`
	SubmittedAt  *string  `json:"submitted_at"`
}

type gradeAssignment struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	PointsPossible float64 `json:"points_possible"`
}

func newGradeExportCmd(flags *rootFlags) *cobra.Command {
	var courseID string
	var outputPath string
	cmd := &cobra.Command{
		Use:   "grade-export",
		Short: "Export student grades for a course as CSV",
		RunE: func(cmd *cobra.Command, args []string) error {
			cc, err := newCanvasClientFromConfig(flags.configPath)
			if err != nil {
				return err
			}
			w := cmd.OutOrStdout()
			if outputPath != "" {
				f, err := os.Create(outputPath)
				if err != nil {
					return fmt.Errorf("creating output file: %w", err)
				}
				defer f.Close()
				w = f
			}
			return runGradeExport(cc, courseID, w)
		},
	}
	cmd.Flags().StringVarP(&courseID, "course", "c", "", "Course ID (required)")
	cmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output file path (default: stdout)")
	_ = cmd.MarkFlagRequired("course")
	return cmd
}

func runGradeExport(cc *canvasClient, courseID string, w io.Writer) error {
	// Fetch enrollments to get student list (students only)
	enrollRaw, err := cc.getAll(fmt.Sprintf("/courses/%s/enrollments", courseID), url.Values{"type[]": {"StudentEnrollment"}})
	if err != nil {
		return fmt.Errorf("fetching enrollments: %w", err)
	}
	type enrollEntry struct {
		UserID int
		Name   string
	}
	var students []enrollEntry
	for _, raw := range enrollRaw {
		var e studentEnrollment
		if err := json.Unmarshal(raw, &e); err != nil || e.UserID == 0 {
			continue
		}
		students = append(students, enrollEntry{UserID: e.UserID, Name: e.User.Name})
	}

	// Fetch assignments
	assignRaw, err := cc.getAll(fmt.Sprintf("/courses/%s/assignments", courseID), nil)
	if err != nil {
		return fmt.Errorf("fetching assignments: %w", err)
	}
	var assignments []gradeAssignment
	for _, raw := range assignRaw {
		var a gradeAssignment
		if err := json.Unmarshal(raw, &a); err != nil || a.ID == 0 {
			continue
		}
		assignments = append(assignments, a)
	}

	// Fetch submissions and index by (user_id, assignment_id)
	subRaw, err := cc.getAll(fmt.Sprintf("/courses/%s/students/submissions", courseID), nil)
	if err != nil {
		return fmt.Errorf("fetching submissions: %w", err)
	}
	type subKey struct{ UserID, AssignmentID int }
	subMap := map[subKey]gradeSubmission{}
	for _, raw := range subRaw {
		var s gradeSubmission
		if err := json.Unmarshal(raw, &s); err != nil {
			continue
		}
		subMap[subKey{s.UserID, s.AssignmentID}] = s
	}

	// Write CSV
	cw := csv.NewWriter(w)
	_ = cw.Write([]string{"student_name", "student_id", "assignment_name", "assignment_id", "score", "possible", "submitted_at"})
	for _, stu := range students {
		for _, a := range assignments {
			sub := subMap[subKey{stu.UserID, a.ID}]
			score := ""
			if sub.Score != nil {
				score = strconv.FormatFloat(*sub.Score, 'f', -1, 64)
			}
			possible := strconv.FormatFloat(a.PointsPossible, 'f', -1, 64)
			submittedAt := ""
			if sub.SubmittedAt != nil {
				submittedAt = *sub.SubmittedAt
			}
			_ = cw.Write([]string{
				stu.Name,
				strconv.Itoa(stu.UserID),
				a.Name,
				strconv.Itoa(a.ID),
				score,
				possible,
				submittedAt,
			})
		}
	}
	cw.Flush()
	return cw.Error()
}
