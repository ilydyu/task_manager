package service

import (
	"context"
	"fmt"

	"github.com/alexedwards/argon2id"
	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/internal/dto"
	"github.com/ilydyu/task_manager.git/pkg/auth"
)

func (s *Service) CreateUser(ctx context.Context, input dto.CreateUserInput) (dto.CreateUserOutput, error) {
	var output dto.CreateUserOutput

	hash, err := argon2id.CreateHash(input.Password, argon2id.DefaultParams)
	if err != nil {
		return output, fmt.Errorf("argon2id.CreateHash: %w", err)
	}

	user, err := domain.NewUser(input.Name, input.Email, hash)

	if err != nil {
		return output, fmt.Errorf("domain.NewUser: %w", err)
	}

	err = s.repository.CreateUser(ctx, &user)

	if err != nil {
		return output, fmt.Errorf("s.repository.CreateUser: %w", err)
	}

	token, err := auth.GenerateToken(user.ID, s.secret)

	if err != nil {
		return output, fmt.Errorf("auth.GenerateToken: %w", err)
	}

	return dto.CreateUserOutput{User: user, Token: token}, nil
}
