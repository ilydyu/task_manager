package integration

import (
	"github.com/ilydyu/task_manager.git/internal/domain"
)

func (s *Suite) TestCreateTeam() {
	res, err := s.client.CreateUser(ctx, "John_Create", "old_john@gmail.com", "12345678")

	s.NoError(err)
	s.Equal("John_Create", res.Name)
	s.Equal("old_john@gmail.com", res.Email)
	s.NotNil(res.Token)

	team, err := s.client.CreateTeam(ctx, res.ID, "Best team", res.Token)

	s.NoError(err)
	s.Equal("Best team", team.Name)
	s.Equal(domain.RoleOwner, team.Members[0].Role)
}
