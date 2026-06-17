package service

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (s *Service) GetTaskHistory(ctx context.Context, taskID int64) (dto.GetTaskHistoryOutput, error) {
	history, err := s.repository.GetTaskHistory(ctx, taskID)

	if err != nil {
		return dto.GetTaskHistoryOutput{}, fmt.Errorf("s.repository.GetTaskHistory: %w", err)
	}

	return dto.GetTaskHistoryOutput{History: history}, nil
}
