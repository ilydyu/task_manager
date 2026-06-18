package repository

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) GetTopCreators(ctx context.Context) ([]domain.TopCreatorStats, error) {
	const query = `
        with task_stats as (
            select 
                t.team_id,
                t.created_by,
                u.name as user_name,
                count(*) as tasks_created,
                rank() over (
                    partition by t.team_id 
                    order by count(*) desc
                ) as user_rank
            from tasks t
            inner join users u on u.id = t.created_by
            where t.created_at >= date_sub(now(), interval 1 month)
            group by t.team_id, t.created_by, u.name
        )
        select 
            ts.team_id,
            tm.name as team_name,
            ts.created_by as user_id,
            ts.user_name,
            ts.tasks_created,
            ts.user_rank
        from task_stats ts
        inner join teams tm on tm.id = ts.team_id
        where ts.user_rank <= ?
        order by ts.team_id, ts.user_rank
    `
	var stats []domain.TopCreatorStats
	tx := transaction.TryExtractTX(ctx)
	rows, err := tx.QueryContext(ctx, query, 3)

	if err != nil {
		return stats, fmt.Errorf("repository: GetTopCreators: tx.QueryContext: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var s domain.TopCreatorStats

		err := rows.Scan(
			&s.TeamID,
			&s.TeamName,
			&s.UserID,
			&s.UserName,
			&s.TasksCreated,
			&s.UserRank,
		)

		if err != nil {
			return stats, fmt.Errorf("repository: GetTopCreators: rows.Scan: %w", err)
		}

		stats = append(stats, s)
	}

	if rows.Err() != nil {
		return stats, fmt.Errorf("repository: GetTopCreators: rows.Err(): %w", err)
	}

	return stats, nil
}
