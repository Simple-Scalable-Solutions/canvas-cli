package cli

import "github.com/spf13/cobra"

// registerTranscendenceCommands adds the hand-crafted compound commands to the root.
// Called from newRootCmd in root.go.
func registerTranscendenceCommands(root *cobra.Command, flags *rootFlags) {
	root.AddCommand(newStudentPulseCmd(flags))
	root.AddCommand(newCourseAuditCmd(flags))
	// Tasks 9-11 will add: newGradeExportCmd, newBulkEnrollCmd, newActivityReportCmd
}
