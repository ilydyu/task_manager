package dto

import "github.com/ilydyu/task_manager.git/internal/domain"

type GetTasksOutput struct {
	Tasks []domain.Task `json:"tasks"`
}
