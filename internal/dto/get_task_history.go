package dto

import "github.com/ilydyu/task_manager.git/internal/domain"

type GetTaskHistoryOutput struct {
	History []domain.TaskHistory
}
