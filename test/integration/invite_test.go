package integration

import (
	"testing"

	"github.com/ilydyu/task_manager.git/internal/domain"
)

func (s *Suite) TestInvite() {
	user, err := s.client.CreateUser(ctx, "Olga_Create", "olga@gmail.com", "12345678")

	s.NoError(err)
	s.Equal("Olga_Create", user.Name)
	s.Equal("olga@gmail.com", user.Email)
	s.NotNil(user.Token)

	team, err := s.client.CreateTeam(ctx, user.ID, "Best team", user.Token)

	s.NoError(err)
	s.Equal("Best team", team.Name)
	s.Equal(domain.RoleOwner, team.Members[0].Role)

	teams, err := s.client.GetUserTeams(ctx, user.Token)

	s.NoError(err)
	s.Equal(1, len(teams.Teams))

	anotherUser, err := s.client.CreateUser(ctx, "Mihai_Create", "Mihai@gmail.com", "12345678")

	s.NoError(err)
	s.Equal("Mihai_Create", anotherUser.Name)
	s.Equal("Mihai@gmail.com", anotherUser.Email)
	s.NotNil(anotherUser.Token)

	s.T().Run("owner invite", func(t *testing.T) {
		res, err := s.client.Invite(ctx, user.ID, teams.Teams[0].ID, anotherUser.ID, user.Token)

		s.NoError(err)
		s.Equal(domain.RoleMember, res.Role)
		s.Equal(anotherUser.ID, res.UserID)
		s.Equal(teams.Teams[0].ID, res.TeamID)
	})

	s.T().Run("member invite", func(t *testing.T) {
		_, err := s.client.Invite(ctx, anotherUser.ID, teams.Teams[0].ID, user.ID, anotherUser.Token)

		s.Error(err)
	})
}
