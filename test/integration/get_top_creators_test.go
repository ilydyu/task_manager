package integration

import (
	"fmt"
	"time"
)

func (s *Suite) TestGetTopCreators() {
	user1, err := s.client.CreateUser(ctx, "Olga", "olga@gmail.com", "12345678")

	s.NoError(err)
	user2, err := s.client.CreateUser(ctx, "Ivan", "Ivan@gmail.com", "12345678")

	s.NoError(err)

	user3, err := s.client.CreateUser(ctx, "Mike", "Mike@gmail.com", "12345678")

	s.NoError(err)

	user4, err := s.client.CreateUser(ctx, "Mihai", "Mihai@gmail.com", "12345678")

	s.NoError(err)

	team1, err := s.client.CreateTeam(ctx, user1.ID, "Olga team", user1.Token)

	s.NoError(err)

	team2, err := s.client.CreateTeam(ctx, user2.ID, "Ivan team", user2.Token)

	s.NoError(err)

	_, err = s.client.Invite(ctx, user1.ID, team1.ID, user2.ID, user1.Token)

	s.NoError(err)

	_, err = s.client.Invite(ctx, user1.ID, team1.ID, user3.ID, user1.Token)

	s.NoError(err)

	_, err = s.client.Invite(ctx, user1.ID, team1.ID, user4.ID, user1.Token)

	s.NoError(err)

	_, err = s.client.Invite(ctx, user2.ID, team2.ID, user1.ID, user2.Token)

	s.NoError(err)

	_, err = s.client.Invite(ctx, user2.ID, team2.ID, user3.ID, user2.Token)

	s.NoError(err)

	_, err = s.client.Invite(ctx, user2.ID, team2.ID, user4.ID, user2.Token)

	s.NoError(err)

	for i := range 3 {
		deadline := time.Now().Add(24 * time.Hour).Truncate(time.Second)
		_, err := s.client.CreateTask(ctx, team1.ID, user1.ID, user1.ID, fmt.Sprintf("task %d", i), fmt.Sprintf("desc %d", i), "done", "low", user1.Token, deadline)

		s.NoError(err)

		_, err = s.client.CreateTask(ctx, team1.ID, user4.ID, user4.ID, fmt.Sprintf("task %d", i), fmt.Sprintf("desc %d", i), "done", "low", user4.Token, deadline)

		s.NoError(err)
	}

	for i := range 2 {
		deadline := time.Now().Add(24 * time.Hour).Truncate(time.Second)
		_, err := s.client.CreateTask(ctx, team2.ID, user2.ID, user2.ID, fmt.Sprintf("task %d", i), fmt.Sprintf("desc %d", i), "done", "low", user2.Token, deadline)

		s.NoError(err)

		_, err = s.client.CreateTask(ctx, team2.ID, user1.ID, user1.ID, fmt.Sprintf("task %d", i), fmt.Sprintf("desc %d", i), "done", "low", user1.Token, deadline)

		s.NoError(err)
	}

	for i := range 1 {
		deadline := time.Now().Add(24 * time.Hour).Truncate(time.Second)
		_, err = s.client.CreateTask(ctx, team1.ID, user3.ID, user3.ID, fmt.Sprintf("task %d", i), fmt.Sprintf("desc %d", i), "done", "low", user3.Token, deadline)

		s.NoError(err)

		_, err = s.client.CreateTask(ctx, team1.ID, user1.ID, user1.ID, fmt.Sprintf("task %d", i), fmt.Sprintf("desc %d", i), "done", "low", user1.Token, deadline)

		s.NoError(err)

		_, err = s.client.CreateTask(ctx, team2.ID, user3.ID, user3.ID, fmt.Sprintf("task %d", i), fmt.Sprintf("desc %d", i), "done", "low", user3.Token, deadline)

		s.NoError(err)

		_, err := s.client.CreateTask(ctx, team2.ID, user2.ID, user2.ID, fmt.Sprintf("task %d", i), fmt.Sprintf("desc %d", i), "done", "low", user2.Token, deadline)

		s.NoError(err)
	}

	stats, err := s.client.GetTopCreators(ctx, user1.Token)

	s.NoError(err)

	for _, t := range stats.Data {
		s.Equal(3, len(t.TopUsers))
		if t.TeamName == "Olga team" {
			s.Equal("Olga", t.TopUsers[0].UserName)
			s.Equal(1, t.TopUsers[0].Rank)
			s.Equal("Mihai", t.TopUsers[1].UserName)
			s.Equal(2, t.TopUsers[1].Rank)
		} else {
			s.Equal("Ivan", t.TopUsers[0].UserName)
			s.Equal(1, t.TopUsers[0].Rank)
			s.Equal("Olga", t.TopUsers[1].UserName)
			s.Equal(2, t.TopUsers[1].Rank)
		}
	}
}
