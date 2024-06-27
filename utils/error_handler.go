package utils

import (
	"errors"
)

var (
	ErrDuplicateUser = errors.New("user exists (duplicate)")
	ErrUserNotFound  = errors.New("user not found")
	ErrQuerying      = errors.New("query exec failed")
)
