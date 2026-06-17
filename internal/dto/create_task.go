package dto

import (
	"time"

	"github.com/ilydyu/task_manager.git/internal/domain"
)

type CreateTaskInput struct {
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`
	TeamID      int64      `json:"team_id"`
	AssigneeID  *int64     `json:"assignee_id,omitempty"`
	CreatedBy   int64      `json:"created_by"`
	Deadline    *time.Time `json:"deadline,omitempty"`
}

type CreateTaskOutput struct {
	domain.Task
}
