package service

import (
	"context"
	"fmt"

	"github.com/alexedwards/argon2id"
	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/internal/dto"
	"github.com/ilydyu/task_manager.git/pkg/auth"
)

func (s *Service) Login(ctx context.Context, input dto.LoginInput) (dto.LoginOutput, error) {
	var output dto.LoginOutput

	user, err := s.repository.GetUserByEmail(ctx, input.Email)

	if err != nil {
		return output, fmt.Errorf("s.repository.GetUserByEmail: %w", err)
	}

	match, err := argon2id.ComparePasswordAndHash(input.Password, user.PasswordHash)
	if err != nil {
		return output, fmt.Errorf("argon2id.CreateHash: %w", err)
	}

	if !match {
		return output, domain.ErrInvalidEmailPassword
	}

	token, err := auth.GenerateToken(user.ID, s.secret)

	if err != nil {
		return output, fmt.Errorf("auth.GenerateToken: %w", err)
	}

	return dto.LoginOutput{User: user, Token: token}, nil
}
