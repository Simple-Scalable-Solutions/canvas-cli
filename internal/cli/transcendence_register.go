package cli

import "github.com/spf13/cobra"

// registerTranscendenceCommands adds the hand-crafted compound commands to the root.
// Called from newRootCmd in root.go.
func registerTranscendenceCommands(root *cobra.Command, flags *rootFlags) {
	root.AddCommand(newStudentPulseCmd(flags))
	// Tasks 8-11 will add: newCourseAuditCmd, newGradeExportCmd, newBulkEnrollCmd, newActivityReportCmd
}
