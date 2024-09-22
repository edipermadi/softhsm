package services

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) GenerateKey(request *messages.GenerateKeyRequest) (*messages.GenerateKeyResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) GenerateKeyPair(request *messages.GenerateKeyPairRequest) (*messages.GenerateKeyPairResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) WrapKey(request *messages.WrapKeyRequest) (*messages.WrapKeyResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) UnwrapKey(request *messages.UnwrapKeyRequest) (*messages.UnwrapKeyResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) DeriveKey(request *messages.DeriveKeyRequest) (*messages.DeriveKeyResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
