package service

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (s *Service) GetInvalidAssignments(ctx context.Context) (dto.GetInvalidAssignmentsOutput, error) {
	var output dto.GetInvalidAssignmentsOutput

	tasks, err := s.repository.GetInvalidAssignments(ctx)

	if err != nil {
		return output, fmt.Errorf("s.repository.GetInvalidAssignments: %w", err)
	}

	return dto.GetInvalidAssignmentsOutput{Data: tasks}, nil
}
