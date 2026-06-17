package service

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/internal/dto"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (s *Service) CreateTask(ctx context.Context, input dto.CreateTaskInput) (dto.CreateTaskOutput, error) {
	var output dto.CreateTaskOutput

	exists, err := s.repository.IsTeamMemberExists(ctx, input.CreatedBy, input.TeamID)

	if err != nil {
		return output, fmt.Errorf("s.repository.IsTeamMemberExists: %w", err)
	}

	if !exists {
		return output, domain.ErrNotAllowed
	}

	task, err := domain.NewTask(
		input.Title,
		input.Description,
		domain.TaskStatus(input.Status),
		domain.TaskPriority(input.Priority),
		input.TeamID,
		input.CreatedBy,
		input.AssigneeID,
		input.Deadline,
	)

	if err != nil {
		return output, fmt.Errorf("domain.NewTask: %w", err)
	}

	err = transaction.Wrap(ctx, func(ctx context.Context) error {
		err = s.repository.CreateTask(ctx, &task)

		if err != nil {
			return fmt.Errorf("s.repository.CreateTask: %w", err)
		}

		err = s.repository.TrackTaskChanges(ctx, task, "create", nil)

		if err != nil {
			return fmt.Errorf("s.repository.TrackTaskChanges: %w", err)
		}

		return nil
	})

	if err != nil {
		return output, fmt.Errorf("transaction.Wrap: %w", err)
	}

	return dto.CreateTaskOutput{Task: task}, nil
}
