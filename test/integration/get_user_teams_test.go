package integration

import (
	"github.com/ilydyu/task_manager.git/internal/domain"
)

func (s *Suite) TestGetUserTeams() {
	res, err := s.client.CreateUser(ctx, "Mike_Create", "mike@gmail.com", "12345678")

	s.NoError(err)
	s.Equal("Mike_Create", res.Name)
	s.Equal("mike@gmail.com", res.Email)
	s.NotNil(res.Token)

	task, err := s.client.CreateTeam(ctx, res.ID, "Best team", res.Token)

	s.NoError(err)
	s.Equal("Best team", task.Name)
	s.Equal(domain.RoleOwner, task.Members[0].Role)

	teams, err := s.client.GetUserTeams(ctx, res.Token)

	s.NoError(err)
	s.Equal(1, len(teams.Teams))
}
