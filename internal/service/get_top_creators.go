package service

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (s *Service) GetTopCreators(ctx context.Context) (dto.GetTopCreatorsOutput, error) {
	var output dto.GetTopCreatorsOutput

	stats, err := s.repository.GetTopCreators(ctx)

	if err != nil {
		return output, fmt.Errorf("s.repository.GetTopCreators: %w", err)
	}

	teamMap := make(map[int64]*dto.TeamTopCreators)
	for _, s := range stats {
		if _, exists := teamMap[s.TeamID]; !exists {
			teamMap[s.TeamID] = &dto.TeamTopCreators{
				TeamID:   s.TeamID,
				TeamName: s.TeamName,
				TopUsers: []dto.TopUserStats{},
			}
		}

		teamMap[s.TeamID].TopUsers = append(teamMap[s.TeamID].TopUsers, dto.TopUserStats{
			UserID:       s.UserID,
			UserName:     s.UserName,
			TasksCreated: s.TasksCreated,
			Rank:         s.UserRank,
		})
	}

	data := dto.GetTopCreatorsOutput{
		Data: make([]dto.TeamTopCreators, 0, len(teamMap)),
	}

	for _, team := range teamMap {
		data.Data = append(data.Data, *team)
	}

	return data, nil
}
