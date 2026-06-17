package domain

import (
	"fmt"
	"time"
)

type TaskStatus string

const (
	TaskStatusBacklog    TaskStatus = "backlog"
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusReview     TaskStatus = "review"
	TaskStatusDone       TaskStatus = "done"
	TaskStatusCancelled  TaskStatus = "cancelled"
)

type TaskPriority string

const (
	TaskPriorityLow    TaskPriority = "low"
	TaskPriorityMedium TaskPriority = "medium"
	TaskPriorityHigh   TaskPriority = "high"
	TaskPriorityUrgent TaskPriority = "urgent"
)

type Task struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title" validate:"required,min=2,max=500"`
	Description *string      `json:"description"`
	Status      TaskStatus   `json:"status" validate:"required,oneof=backlog todo in_progress review done cancelled"`
	Priority    TaskPriority `json:"priority" validate:"required,oneof=low medium high urgent"`
	TeamID      int64        `json:"team_id" validate:"required"`
	AssigneeID  *int64       `json:"assignee_id"`
	CreatedBy   int64        `json:"created_by" validate:"required"`
	Deadline    *time.Time   `json:"deadline"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func NewTask(title string, description *string, status TaskStatus, priority TaskPriority, teamID, createdByID int64, assigneeID *int64, deadline *time.Time) (Task, error) {
	t := Task{
		Title:       title,
		Description: description,
		Status:      status,
		Priority:    priority,
		TeamID:      teamID,
		AssigneeID:  assigneeID,
		CreatedBy:   createdByID,
		Deadline:    deadline,
	}

	if err := t.Validate(); err != nil {
		return Task{}, fmt.Errorf("p.Validate: %w", err)
	}

	return t, nil
}

func (t Task) Validate() error {
	err := validate.Struct(t)

	if err != nil {
		return fmt.Errorf("validate.Struct: %w", err)
	}

	return nil
}
