package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/internal/dto"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (s *Service) UpdateTask(ctx context.Context, input dto.UpdateTaskInput, userID, taskID int64) (dto.UpdateTaskOutput, error) {
	var output dto.UpdateTaskOutput

	task, err := s.repository.GetTaskByID(ctx, taskID)

	if err != nil {
		return output, fmt.Errorf("s.repository.GetTaskByID: %w", err)
	}

	member, err := s.repository.GetTeamMember(ctx, userID, task.TeamID)

	if err != nil {
		return output, fmt.Errorf("s.repository.GetTeamMember: %w", err)
	}

	if member.Role != domain.RoleOwner && member.Role != domain.RoleAdmin {
		return output, domain.ErrNotAllowed
	}

	newTask, err := domain.NewTask(
		input.Title,
		input.Description,
		domain.TaskStatus(input.Status),
		domain.TaskPriority(input.Priority),
		task.TeamID,
		task.CreatedBy,
		input.AssigneeID,
		input.Deadline,
	)

	if err != nil {
		return output, fmt.Errorf("domain.NewTask: %w", err)
	}

	oldValue := newTask
	newTask.ID = taskID

	err = transaction.Wrap(ctx, func(ctx context.Context) error {
		err = s.repository.UpdateTask(ctx, &newTask)

		if err != nil {
			return fmt.Errorf("s.repository.UpdateTask: %w", err)
		}

		data, err := json.Marshal(oldValue)

		if err != nil {
			return fmt.Errorf("service: UpdateTask: json.Marshal: %w", err)
		}

		err = s.repository.TrackTaskChanges(ctx, newTask, "update", data)

		if err != nil {
			return fmt.Errorf("s.repository.TrackTaskChanges: %w", err)
		}

		return nil
	})

	if err != nil {
		return output, fmt.Errorf("transaction.Wrap: %w", err)
	}

	return dto.UpdateTaskOutput{Task: newTask}, nil
}
