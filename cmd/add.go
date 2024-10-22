package cmd

import (
	"github.com/spf13/cobra"
)

func addCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a new task to your list",
		Run: func(cmd *cobra.Command, args []string) {
			// Your code here
		},
	}
}