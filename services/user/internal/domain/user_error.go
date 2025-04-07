package domain

import "errors"

var (
	ErrUserAlreadyExists = errors.New("email already exists")
)
