package cmd

import (
	"errors"

	"github.com/livanjimenez/task-tracker-cli/internal/task"
	"github.com/spf13/cobra"
)

func AddCmd() *cobra.Command {
	cmd := &cobra.Command {
		Use: "add",
		Short: "Add a new task to the list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddTaskCmd(args)
		},
	}
	return cmd
}

func RunAddTaskCmd(args []string) error {
	if len(args) == 0 {
		return errors.New("Task description is required")
	}

	description := args[0]
	return task.AddTask(description)

}