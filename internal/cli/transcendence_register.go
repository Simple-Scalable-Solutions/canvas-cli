package cli

import "github.com/spf13/cobra"

// registerTranscendenceCommands adds the hand-crafted compound commands to the root.
// Called from newRootCmd in root.go.
func registerTranscendenceCommands(root *cobra.Command, flags *rootFlags) {
	root.AddCommand(newStudentPulseCmd(flags))
	root.AddCommand(newCourseAuditCmd(flags))
	root.AddCommand(newGradeExportCmd(flags))
	root.AddCommand(newBulkEnrollCmd(flags))
	// Task 11 will add: newActivityReportCmd
}
