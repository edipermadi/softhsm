package services

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) SeedRandom(request *messages.SeedRandomRequest) (*messages.SeedRandomResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) GenerateRandom(request *messages.GenerateRandomRequest) (*messages.GenerateRandomResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
