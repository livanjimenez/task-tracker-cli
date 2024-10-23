package task

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

type TaskStatus string

const (
	TASK_STATUS_TODO 				TaskStatus = "TODO"
	TASK_STATUS_DONE 				TaskStatus = "DONE"
	TASK_STATUS_IN_PROGRESS TaskStatus = "IN_PROGRESS"
)

type Task struct {
	ID	 				int64				`json:"id"`
	Description string			`json:"description"`
	Status			TaskStatus	`json:"status"`
	CreatedAt		time.Time		`json:"created_at"`
	UpdatedAt		time.Time		`json:"updated_at"`
}

func NewTask(id int64, description string) *Task {
	return &Task{
		ID:						id,
		Description: 	description,
		Status:				TASK_STATUS_TODO,
		CreatedAt:		time.Now(),
		UpdatedAt:		time.Now(),
	}
}

func AddTask(description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var newTaskId int64
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newTaskId = lastTask.ID + 1
	} else {
		newTaskId = 1
	}

	task := NewTask(newTaskId, description)
	tasks = append(tasks, *task)

	style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFCC66"))
	formattedId := style.Render(fmt.Sprintf("(ID: %d)", task.ID))
	fmt.Printf("Task added: %s %s\n", task.Description, formattedId)

	return WriteTasksToFile(tasks)
}

func UpdateTaskDescription(id int64, description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var taskExists bool = false
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			taskExists = true
			task.Description = description
			task.UpdatedAt = time.Now()
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !taskExists {
		return fmt.Errorf("task not found (ID: %d)", id)
	}

	formattedId := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66")).
		Render(fmt.Sprintf("(ID: %d)", id))
	fmt.Printf("\nTask updated successfully: %s\n\n", formattedId)
	return WriteTasksToFile(updatedTasks)
}

func DeleteTask(id int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var taskExists bool = false
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			taskExists = true
			continue
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !taskExists {
		return fmt.Errorf("task not found (ID: %d)", id)
	}

	formattedId := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66")).
		Render(fmt.Sprintf("(ID: %d)", id))
	fmt.Printf("\nTask deleted successfully: %s\n\n", formattedId)
	return WriteTasksToFile(updatedTasks)
}

func ListTasks(status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var tasksToDisplay []Task
	for _, task := range tasks {
		if status == "all" || task.Status == status {
			tasksToDisplay = append(tasksToDisplay, task)
		}
	}

	if len(tasksToDisplay) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	fmt.Println()
	for _, task := range tasksToDisplay {
		formattedId := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFCC66")).
			Render(fmt.Sprintf("(ID: %d)", task.ID))
		fmt.Printf("%s %s\n", task.Description, formattedId)
	}
	fmt.Println()

	return nil
}
