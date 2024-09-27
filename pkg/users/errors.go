package users

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrPasswordIsExpired   = errors.New("password is expired")
	ErrPasswordIsIncorrect = errors.New("password is incorrect")
	ErrPasswordIsNotSet    = errors.New("password is not set")
)
