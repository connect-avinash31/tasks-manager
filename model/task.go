// package model provides the implementation of the task management service.
package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TaskStatus string

// TaskStatus represents the status of a task in the task manager system.
// It can be "Todo", "InProgress", "Done", etc.
var (
	Todo       TaskStatus = "Todo"
	InProgress TaskStatus = "InProgress"
	Done       TaskStatus = "Done"
)

// add the json tags to the TaskStatus type
func (s TaskStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + string(s) + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for TaskStatus.
func (s *TaskStatus) UnmarshalJSON(data []byte) error {
	var status string
	if err := json.Unmarshal(data, &status); err != nil {
		return fmt.Errorf("invalid task status: %w", err)
	}
	// convert in lowercase to make it case-insensitive
	status = strings.ToLower(status)
	switch status {
	case strings.ToLower(string(Todo)):
		*s = Todo
	case strings.ToLower(string(InProgress)):
		*s = InProgress
	case strings.ToLower(string(Done)):
		*s = Done
	default:
		return fmt.Errorf("unknown task status: %s", status)
	}
	return nil
}

// Task represents a job/task in task-manager syste, , it has details like
// title , description , status etc.
type Task struct {
	Title       string     `json:"title" validate:"required" `   // title of the task
	Description string     `json:"description"`                  // detailed description of the task
	DueDate     string     `json:"due_date" validate:"required"` // date like "2023-10-01"
	Status      TaskStatus `json:"status" validate:"required"`   // e.g., "Todo", "InProgress", "Done"
}
