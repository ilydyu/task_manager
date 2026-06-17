package repository

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) GetUserTeams(ctx context.Context, userID int) ([]domain.Team, error) {
	const query = `select distinct teams.* from teams join team_members on teams.id = team_members.team_id where team_members.user_id = ?`

	tx := transaction.TryExtractTX(ctx)

	var teams []domain.Team

	rows, err := tx.QueryContext(ctx, query, userID)

	if err != nil {
		return teams, fmt.Errorf("tx.QueryContext: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var team domain.Team

		err := rows.Scan(&team.ID, &team.Name, &team.CreatedBy, &team.CreatedAt, &team.UpdatedAt)

		if err != nil {
			return teams, fmt.Errorf("rows.Scan: %w", err)
		}

		teams = append(teams, team)
	}

	if rows.Err() != nil {
		return teams, fmt.Errorf("rows.Err: %w", err)
	}

	return teams, nil
}
