package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name" validate:"required,min=3,max=255"`
	Email        string    `json:"email" validate:"required,email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

var validate = validator.New(validator.WithRequiredStructEnabled()) //nolint:gochecknoglobals

func NewUser(name, email, password_hash string) (User, error) {
	u := User{
		Name:         name,
		Email:        email,
		PasswordHash: password_hash,
	}

	if err := u.Validate(); err != nil {
		return User{}, fmt.Errorf("p.Validate: %w", err)
	}

	return u, nil
}

func (u User) Validate() error {
	err := validate.Struct(u)
	if err != nil {
		return fmt.Errorf("validate.Struct: %w", err)
	}

	return nil
}
