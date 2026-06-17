package dto

import (
	"time"

	"github.com/ilydyu/task_manager.git/internal/domain"
)

type UpdateTaskInput struct {
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`
	AssigneeID  *int64     `json:"assignee_id"`
	Deadline    *time.Time `json:"deadline"`
}

type UpdateTaskOutput struct {
	domain.Task
}
