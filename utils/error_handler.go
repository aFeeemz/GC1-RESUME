package utils

import (
	"errors"
)

var (
	ErrDuplicateUser = errors.New("user exists (duplicate)")
	ErrUserNotFound  = errors.New("user not found")
	ErrFieldNotFound = errors.New("field not found")
	ErrQuerying      = errors.New("query exec failed")
	ErrGetData       = errors.New("get data failed")
)
