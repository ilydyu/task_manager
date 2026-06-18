package integration

import (
	"fmt"
	"time"
)

func (s *Suite) TestGetTasks() {
	user, err := s.client.CreateUser(ctx, "Olga_Create", "olga@gmail.com", "12345678")

	s.NoError(err)

	for i := range 4 {
		team, err := s.client.CreateTeam(ctx, user.ID, fmt.Sprintf("Best team %d", i), user.Token)

		s.NoError(err)

		for j := range 10 {
			deadline := time.Now().Add(24 * time.Hour).Truncate(time.Second)
			_, err := s.client.CreateTask(ctx, team.ID, user.ID, user.ID, fmt.Sprintf("task %d", j), fmt.Sprintf("desc %d", j), "done", "low", user.Token, deadline)

			s.NoError(err)
		}

		deadline := time.Now().Add(24 * time.Hour).Truncate(time.Second)
		_, err = s.client.CreateTask(ctx, team.ID, user.ID, user.ID, fmt.Sprintf("task by team %d", i), fmt.Sprintf("desc team%d", i), "backlog", "low", user.Token, deadline)

		s.NoError(err)
	}

	stats, err := s.client.GetTeamStats(ctx, user.Token)

	s.NoError(err)
	s.Equal(4, len(stats.Data))

	for _, t := range stats.Data {
		s.Equal(1, t.MemberCount)
		s.Equal(10, t.DoneTasksLast7Days)
	}
}
