package services

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) SignInit(request *messages.SignInitRequest) (*messages.SignInitResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) Sign(request *messages.SignRequest) (*messages.SignResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SignUpdate(request *messages.SignUpdateRequest) (*messages.SignUpdateResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SignFinal(request *messages.SignFinalRequest) (*messages.SignFinalResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SignRecoverInit(request *messages.SignRecoverInitRequest) (*messages.SignRecoverInitResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SignRecover(request *messages.SignRecoverRequest) (*messages.SignRecoverResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
