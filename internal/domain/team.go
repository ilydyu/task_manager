package domain

import (
	"fmt"
	"time"
)

type Team struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=255"`
	CreatedBy int64     `json:"user" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTeam(userID int64, name string) (Team, error) {
	t := Team{
		CreatedBy: userID,
		Name:      name,
	}

	if err := t.Validate(); err != nil {
		return Team{}, fmt.Errorf("p.Validate: %w", err)
	}

	return t, nil
}

func (t Team) Validate() error {
	err := validate.Struct(t)
	if err != nil {
		return fmt.Errorf("validate.Struct: %w", err)
	}

	return nil
}
