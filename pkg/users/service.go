package users

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Authenticate(ctx context.Context, userType Type, username string, password string) error
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

type service struct {
	repository Repository
}

func (s *service) Authenticate(ctx context.Context, userType Type, username string, password string) error {
	user, err := s.repository.GetUser(ctx, username)
	if err != nil {
		return err
	}

	if user.PasswordExpiredAt.Before(time.Now()) {
		return ErrPasswordIsExpired
	}

	var succeeded bool
	defer func() { _ = s.repository.RecordLoginAttempt(ctx, user.ID, succeeded) }()

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return ErrPasswordIsIncorrect
	}

	if errors.Is(err, bcrypt.ErrHashTooShort) {
		return ErrPasswordIsNotSet
	}

	succeeded = true
	return nil
}
