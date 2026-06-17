package integration

import (
	"testing"
	"time"

	"github.com/ilydyu/task_manager.git/internal/domain"
)

func (s *Suite) TestCreateTask() {
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

	res, err := s.client.Invite(ctx, user.ID, teams.Teams[0].ID, anotherUser.ID, user.Token)

	s.NoError(err)
	s.Equal(domain.RoleMember, res.Role)
	s.Equal(anotherUser.ID, res.UserID)
	s.Equal(teams.Teams[0].ID, res.TeamID)

	thirdUser, err := s.client.CreateUser(ctx, "Ivan_Create", "Ivan@gmail.com", "12345678")

	s.NoError(err)
	s.Equal("Ivan_Create", thirdUser.Name)
	s.Equal("Ivan@gmail.com", thirdUser.Email)
	s.NotNil(thirdUser.Token)

	s.T().Run("member create task", func(t *testing.T) {
		deadline := time.Now().Add(24 * time.Hour).Truncate(time.Second)
		task, err := s.client.CreateTask(ctx, team.ID, anotherUser.ID, user.ID, "task", "desc", "backlog", "low", user.Token, deadline)

		s.NoError(err)
		s.Equal("task", task.Title)
		s.Equal("desc", *task.Description)
		s.Equal(domain.TaskStatusBacklog, task.Status)
		s.Equal(domain.TaskPriorityLow, task.Priority)
		s.Equal(team.ID, task.TeamID)
		s.Equal(anotherUser.ID, *task.AssigneeID)
		s.Equal(user.ID, task.CreatedBy)
		s.Equal(deadline, *task.Deadline)
	})

	s.T().Run("foreign user trying create task", func(t *testing.T) {
		deadline := time.Now().Add(24 * time.Hour).Truncate(time.Second)
		_, err := s.client.CreateTask(ctx, team.ID, anotherUser.ID, thirdUser.ID, "task", "desc", "backlog", "low", thirdUser.Token, deadline)

		s.Error(err)
	})
}
