package dto

import "github.com/ilydyu/task_manager.git/internal/domain"

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserOutput struct {
	domain.User
	Token string `json:"token"`
}
