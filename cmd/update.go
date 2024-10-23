package cmd

import (
	"fmt"
	"strconv"

	"github.com/livanjimenez/task-tracker-cli/internal/task"
	"github.com/spf13/cobra"
)

func UpdateCmd() *cobra.Command {
	cmd := &cobra.Command {
		Use: "update",
		Short: "Update an existing task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskCmd(args)
		},
	}
	return cmd
}

func RunUpdateTaskCmd(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("please provide a task id and the new description")
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		return err
	}

	newDescription := args[1]
	return task.UpdateTaskDescription(taskIDInt, newDescription)
}