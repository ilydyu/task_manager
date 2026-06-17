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

func (c *Client) CreateTask(ctx context.Context, teamID, assigneeID, createdByID int64, title, description, status, priority, token string, deadline time.Time) (dto.CreateTaskOutput, error) {
	const createTask = "api/v1/tasks"

	path := fmt.Sprintf("http://%s/%s", c.host, createTask)

	request := struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		Priority    string    `json:"priority"`
		TeamID      int64     `json:"team_id"  `
		AssigneeID  int64     `json:"assignee_id"  `
		CreatedBy   int64     `json:"created_by"`
		Deadline    time.Time `json:"deadline"`
	}{
		Title:       title,
		Description: description,
		Status:      status,
		Priority:    priority,
		TeamID:      teamID,
		AssigneeID:  assigneeID,
		CreatedBy:   createdByID,
		Deadline:    deadline,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return dto.CreateTaskOutput{}, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewReader(body))
	if err != nil {
		return dto.CreateTaskOutput{}, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.client.Do(req)
	if err != nil {
		return dto.CreateTaskOutput{}, fmt.Errorf("client.Do: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return dto.CreateTaskOutput{}, fmt.Errorf("request failed: status: %s, body:%s", resp.Status, body)
	}

	response := dto.CreateTaskOutput{}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return dto.CreateTaskOutput{}, fmt.Errorf("json.Decode: %w", err)
	}

	return response, nil
}
