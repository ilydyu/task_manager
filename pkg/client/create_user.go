package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (c *Client) CreateUser(ctx context.Context, name string, email, password string) (dto.CreateUserOutput, error) {
	const createUser = "api/v1/register"

	path := fmt.Sprintf("http://%s/%s", c.host, createUser)

	request := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Name:     name,
		Email:    email,
		Password: password,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return dto.CreateUserOutput{}, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewReader(body))
	if err != nil {
		return dto.CreateUserOutput{}, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return dto.CreateUserOutput{}, fmt.Errorf("client.Do: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return dto.CreateUserOutput{}, fmt.Errorf("request failed: status: %s, body:%s", resp.Status, body)
	}

	response := dto.CreateUserOutput{}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return dto.CreateUserOutput{}, fmt.Errorf("json.Decode: %w", err)
	}

	return response, nil
}
