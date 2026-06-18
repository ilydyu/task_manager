package integration

import (
	"time"
)

func (s *Suite) TestGetInvalidAssigments() {
	user1, err := s.client.CreateUser(ctx, "Olga", "olga@gmail.com", "12345678")

	s.NoError(err)

	user2, err := s.client.CreateUser(ctx, "Ivan", "Ivan@gmail.com", "12345678")

	s.NoError(err)

	team1, err := s.client.CreateTeam(ctx, user1.ID, "Olga team", user1.Token)

	s.NoError(err)

	deadline := time.Now().Add(24 * time.Hour).Truncate(time.Second)
	_, err = s.client.CreateTask(ctx, team1.ID, user1.ID, user1.ID, "task", "desc", "done", "low", user1.Token, deadline)

	s.NoError(err)

	task, err := s.client.CreateTask(ctx, team1.ID, user2.ID, user1.ID, "task", "desc", "done", "low", user1.Token, deadline)

	s.NoError(err)

	tasks, err := s.client.GetInvalidAssignments(ctx, user1.Token)

	s.NoError(err)

	s.Equal(1, len(tasks.Data))
	s.Equal(task.ID, tasks.Data[0].TaskID)
	s.Equal(user2.ID, tasks.Data[0].AssigneeID)
	s.Equal(user2.Name, tasks.Data[0].AssigneeName)
}
