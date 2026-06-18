package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (c *Client) GetInvalidAssignments(ctx context.Context, token string) (dto.GetInvalidAssignmentsOutput, error) {
	const getInvalidAssigments = "api/v1/tasks/invalid_assignments"

	path := fmt.Sprintf("http://%s/%s", c.host, getInvalidAssigments)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, http.NoBody)
	if err != nil {
		return dto.GetInvalidAssignmentsOutput{}, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.client.Do(req)
	if err != nil {
		return dto.GetInvalidAssignmentsOutput{}, fmt.Errorf("client.Do: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return dto.GetInvalidAssignmentsOutput{}, fmt.Errorf("request failed: status: %s", resp.Status)
	}

	response := dto.GetInvalidAssignmentsOutput{}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return dto.GetInvalidAssignmentsOutput{}, fmt.Errorf("json.Decode: %w", err)
	}

	return response, nil
}
