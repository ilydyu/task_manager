package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (s *Service) GetTasks(ctx context.Context, teamID, assigneeID, status, cursor string) (dto.GetTasksOutput, error) {
	cacheKey := fmt.Sprintf("tasks:%s:%s:%s:%s", teamID, assigneeID, status, cursor)

	cached, err := s.redis.Get(ctx, cacheKey)

	if err == nil {
		cachedStr, ok := cached.(string)
		if !ok {
			return dto.GetTasksOutput{}, fmt.Errorf("cached value is not string")
		}
		var tasks []domain.Task

		err := json.Unmarshal([]byte(cachedStr), &tasks)
		if err != nil {
			return dto.GetTasksOutput{}, fmt.Errorf("failed to unmarshal cached data: %w", err)
		}

		return dto.GetTasksOutput{Tasks: tasks}, nil
	}

	tasks, err := s.repository.GetTasks(ctx, teamID, assigneeID, status, cursor)

	if err != nil {
		return dto.GetTasksOutput{}, fmt.Errorf("s.repository.GetTasks: %w", err)
	}

	data, err := json.Marshal(tasks)

	if err != nil {
		return dto.GetTasksOutput{}, fmt.Errorf("service: GetTasks: json.Marshal: %w", err)
	}

	err = s.redis.Set(ctx, cacheKey, string(data), 5*time.Minute)

	if err != nil {
		return dto.GetTasksOutput{}, fmt.Errorf("service: GetTasks: s.redis.Set: %w", err)
	}

	return dto.GetTasksOutput{Tasks: tasks}, nil
}
