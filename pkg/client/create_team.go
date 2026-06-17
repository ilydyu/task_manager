package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (c *Client) CreateTeam(ctx context.Context, userID int64, name, token string) (dto.CreateTeamOutput, error) {
	const createTeam = "api/v1/teams"

	path := fmt.Sprintf("http://%s/%s", c.host, createTeam)

	request := struct {
		UserID int64  `json:"user_id"`
		Name   string `json:"name"`
	}{
		UserID: userID,
		Name:   name,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return dto.CreateTeamOutput{}, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewReader(body))
	if err != nil {
		return dto.CreateTeamOutput{}, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.client.Do(req)
	if err != nil {
		return dto.CreateTeamOutput{}, fmt.Errorf("client.Do: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return dto.CreateTeamOutput{}, fmt.Errorf("request failed: status: %s, body:%s", resp.Status, body)
	}

	response := dto.CreateTeamOutput{}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return dto.CreateTeamOutput{}, fmt.Errorf("json.Decode: %w", err)
	}

	return response, nil
}
