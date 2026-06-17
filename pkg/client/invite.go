package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (c *Client) Invite(ctx context.Context, memberID, teamID, invitedUserID int64, token string) (dto.InviteOutput, error) {
	const invite = "api/v1/teams"

	path := fmt.Sprintf("http://%s/%s/%d/invite", c.host, invite, teamID)

	request := struct {
		InvitedUserID int64 `json:"invited_user_id"`
	}{
		InvitedUserID: invitedUserID,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return dto.InviteOutput{}, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewReader(body))
	if err != nil {
		return dto.InviteOutput{}, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.client.Do(req)
	if err != nil {
		return dto.InviteOutput{}, fmt.Errorf("client.Do: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return dto.InviteOutput{}, fmt.Errorf("request failed: status: %s, body:%s", resp.Status, body)
	}

	response := dto.InviteOutput{}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return dto.InviteOutput{}, fmt.Errorf("json.Decode: %w", err)
	}

	return response, nil
}
