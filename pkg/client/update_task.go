package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (c *Client) UpdateTask(ctx context.Context, taskID, assigneeID int64, title, description, status, priority, token string, deadline time.Time) (dto.UpdateTaskOutput, error) {
	const updateTask = "api/v1/tasks"

	path := fmt.Sprintf("http://%s/%s/%d", c.host, updateTask, taskID)

	request := struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		Priority    string    `json:"priority"`
		AssigneeID  int64     `json:"assignee_id"  `
		Deadline    time.Time `json:"deadline"`
	}{
		Title:       title,
		Description: description,
		Status:      status,
		Priority:    priority,
		AssigneeID:  assigneeID,
		Deadline:    deadline,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return dto.UpdateTaskOutput{}, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, path, bytes.NewReader(body))
	if err != nil {
		return dto.UpdateTaskOutput{}, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.client.Do(req)
	if err != nil {
		return dto.UpdateTaskOutput{}, fmt.Errorf("client.Do: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return dto.UpdateTaskOutput{}, fmt.Errorf("request failed: status: %s, body:%s", resp.Status, body)
	}

	response := dto.UpdateTaskOutput{}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return dto.UpdateTaskOutput{}, fmt.Errorf("json.Decode: %w", err)
	}

	return response, nil
}
