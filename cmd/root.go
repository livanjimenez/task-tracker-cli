package cmd

import "github.com/spf13/cobra"

func RootCmd() *cobra.Command {
	cmd := &cobra.Command {
		Use: "task-tracker-cli",
		Short: "Task Tracker CLI",
	}
	cmd.AddCommand(AddCmd())
	cmd.AddCommand(UpdateCmd())
	cmd.AddCommand(DeleteCmd())
	cmd.AddCommand(ListCmd())
	return cmd
}