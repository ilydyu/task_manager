package domain

import "errors"

var (
	ErrDuplicate            = errors.New("Duplicate record")
	ErrNotFound             = errors.New("Not found")
	ErrInvalidEmailPassword = errors.New("Invalid email or password")
	ErrNotAllowed           = errors.New("Not allowed")
	ErrInvalidAction        = errors.New("Invalid action")
)
