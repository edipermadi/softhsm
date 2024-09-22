package services

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) EncryptInit(request *messages.EncryptInitRequest) (*messages.EncryptInitResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) Encrypt(request *messages.EncryptRequest) (*messages.EncryptResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) EncryptUpdate(request *messages.EncryptUpdateRequest) (*messages.EncryptUpdateResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) EncryptFinal(request *messages.EncryptFinalRequest) (*messages.EncryptFinalResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
