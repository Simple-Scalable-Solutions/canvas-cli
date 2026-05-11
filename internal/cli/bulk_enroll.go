package cli

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func newBulkEnrollCmd(flags *rootFlags) *cobra.Command {
	var courseID string
	var role string
	var filePath string
	cmd := &cobra.Command{
		Use:   "bulk-enroll",
		Short: "Enroll multiple users into a course from a file or stdin",
		RunE: func(cmd *cobra.Command, args []string) error {
			cc, err := newCanvasClientFromConfig(flags.configPath)
			if err != nil {
				return err
			}
			r := cmd.InOrStdin()
			if filePath != "" {
				f, err := os.Open(filePath)
				if err != nil {
					return fmt.Errorf("opening file: %w", err)
				}
				defer f.Close()
				r = f
			}
			return runBulkEnroll(cc, courseID, role, flags.dryRun, r, cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVarP(&courseID, "course", "c", "", "Course ID (required)")
	cmd.Flags().StringVar(&role, "role", "StudentEnrollment", "Enrollment role type")
	cmd.Flags().StringVarP(&filePath, "file", "f", "", "File with user IDs (one per line; default: stdin)")
	_ = cmd.MarkFlagRequired("course")
	return cmd
}

func runBulkEnroll(cc *canvasClient, courseID, role string, dryRun bool, r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	var enrolled, failed int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || line == "user_id" {
			continue
		}
		// Strip CSV: take first field only
		if idx := strings.Index(line, ","); idx >= 0 {
			line = strings.TrimSpace(line[:idx])
		}
		if line == "" {
			continue
		}
		if dryRun {
			fmt.Fprintf(w, "[dry-run] Would enroll user %s as %s in course %s\n", line, role, courseID)
			enrolled++
			continue
		}
		body := url.Values{
			"enrollment[user_id]":          {line},
			"enrollment[type]":             {role},
			"enrollment[enrollment_state]": {"active"},
		}
		_, err := cc.post(fmt.Sprintf("/courses/%s/enrollments", courseID), body)
		if err != nil {
			fmt.Fprintf(w, "Failed user %s: %v\n", line, err)
			failed++
		} else {
			fmt.Fprintf(w, "Enrolled user %s as %s\n", line, role)
			enrolled++
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("reading input: %w", err)
	}
	fmt.Fprintf(w, "\nSummary: %d enrolled, %d failed\n", enrolled, failed)
	return nil
}
