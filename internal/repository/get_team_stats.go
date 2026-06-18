package repository

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) GetTeamStats(ctx context.Context) ([]domain.TeamStatistics, error) {
	const query = `
        select 
			t.id,
            t.name,
            count(distinct tm.user_id) as member_count,
            coalesce(
                count(
                    case 
                        when tk.status = 'done' 
                        and tk.updated_at >= date_sub(now(), interval 7 day) 
                        then 1 
                    end
                ), 
                0
            ) as done_tasks_last_7_days
        from teams t
        left join team_members tm on tm.team_id = t.id
        left join tasks tk on tk.team_id = t.id
        group by t.id, t.name
        order by t.name
    `

	var stats []domain.TeamStatistics
	tx := transaction.TryExtractTX(ctx)

	rows, err := tx.QueryContext(ctx, query)

	if err != nil {
		return stats, fmt.Errorf("repository: GetTeamStats: tx.QueryContext: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var s domain.TeamStatistics

		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.MemberCount,
			&s.DoneTasksLast7Days,
		)

		if err != nil {
			return stats, fmt.Errorf("repository: GetTeamStats: rows.Scan: %w", err)
		}

		stats = append(stats, s)
	}

	if rows.Err() != nil {
		return stats, fmt.Errorf("repository: GetTeamStats: rows.Err(): %w", err)
	}

	return stats, nil
}
