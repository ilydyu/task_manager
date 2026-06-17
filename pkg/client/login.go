package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (c *Client) Login(ctx context.Context, email, password string) (dto.LoginOutput, error) {
	const login = "api/v1/login"

	path := fmt.Sprintf("http://%s/%s", c.host, login)

	request := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    email,
		Password: password,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return dto.LoginOutput{}, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewReader(data))
	if err != nil {
		return dto.LoginOutput{}, fmt.Errorf("http.NewRequest: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return dto.LoginOutput{}, fmt.Errorf("client.Do: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.LoginOutput{}, fmt.Errorf("io.ReadAll: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return dto.LoginOutput{}, ErrNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return dto.LoginOutput{}, fmt.Errorf("request failed: status: %s, body:%s", resp.Status, body)
	}

	var response dto.LoginOutput

	if err = json.Unmarshal(body, &response); err != nil {
		return dto.LoginOutput{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return response, nil
}
