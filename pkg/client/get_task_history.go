package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (c *Client) GetTaskHistory(ctx context.Context, taskID int64, token string) (dto.GetTaskHistoryOutput, error) {
	const getTaskHistory = "api/v1/tasks"

	path := fmt.Sprintf("http://%s/%s/%d/history", c.host, getTaskHistory, taskID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, http.NoBody)
	if err != nil {
		return dto.GetTaskHistoryOutput{}, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.client.Do(req)
	if err != nil {
		return dto.GetTaskHistoryOutput{}, fmt.Errorf("client.Do: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return dto.GetTaskHistoryOutput{}, fmt.Errorf("request failed: status: %s", resp.Status)
	}

	response := dto.GetTaskHistoryOutput{}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return dto.GetTaskHistoryOutput{}, fmt.Errorf("json.Decode: %w", err)
	}

	return response, nil
}
