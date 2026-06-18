package repository

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) GetInvalidAssignments(ctx context.Context) ([]domain.InvalidAssignment, error) {
	const query = `
        select 
            t.id as task_id,
            t.title,
            t.team_id,
            tm.name as team_name,
            t.assignee_id,
            u.name as assignee_name
        from tasks t
        inner join teams tm on tm.id = t.team_id
        inner join users u on u.id = t.assignee_id
        where t.assignee_id is not null
            and not exists (
                select 1
                from team_members tm_members
                where tm_members.team_id = t.team_id
                    and tm_members.user_id = t.assignee_id
            )
        order by t.team_id, t.id
    `
	var tasks []domain.InvalidAssignment
	tx := transaction.TryExtractTX(ctx)

	rows, err := tx.QueryContext(ctx, query)

	if err != nil {
		return tasks, fmt.Errorf("repository: GetInvalidAssignments: tx.QueryContext: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var task domain.InvalidAssignment

		err := rows.Scan(
			&task.TaskID,
			&task.Title,
			&task.TeamID,
			&task.TeamName,
			&task.AssigneeID,
			&task.AssigneeName,
		)

		if err != nil {
			return tasks, fmt.Errorf("repository: GetInvalidAssignments: rows.Scan: %w", err)
		}

		tasks = append(tasks, task)
	}

	if rows.Err() != nil {
		return tasks, fmt.Errorf("repository: GetInvalidAssignments: rows.Err(): %w", err)
	}

	return tasks, nil
}
