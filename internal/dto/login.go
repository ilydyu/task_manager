package dto

import "github.com/ilydyu/task_manager.git/internal/domain"

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	domain.User
	Token string `json:"token"`
}
