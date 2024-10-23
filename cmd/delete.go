package cmd

import (
	"fmt"
	"strconv"

	"github.com/livanjimenez/task-tracker-cli/internal/task"
	"github.com/spf13/cobra"
)

func DeleteCmd() *cobra.Command {
	cmd := &cobra.Command {
		Use: "delete",
		Short: "Delete an existing task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteTaskCmd(args)
		},
	}
	return cmd
}

func RunDeleteTaskCmd(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a task id")
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		return err
	}

	return task.DeleteTask(taskIDInt)
}

