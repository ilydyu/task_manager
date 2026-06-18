package dto

import "github.com/ilydyu/task_manager.git/internal/domain"

type GetInvalidAssignmentsOutput struct {
	Data []domain.InvalidAssignment `json:"data"`
}
