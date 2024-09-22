package services

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) DecryptInit(request *messages.DecryptInitRequest) (*messages.DecryptInitResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) Decrypt(request *messages.DecryptRequest) (*messages.DecryptResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) DecryptUpdate(request *messages.DecryptUpdateRequest) (*messages.DecryptUpdateResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) DecryptFinal(request *messages.DecryptFinalRequest) (*messages.DecryptFinalResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
