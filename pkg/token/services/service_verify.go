package services

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) VerifyInit(request *messages.VerifyInitRequest) (*messages.VerifyInitResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) Verify(request *messages.VerifyRequest) (*messages.VerifyResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) VerifyUpdate(request *messages.VerifyUpdateRequest) (*messages.VerifyUpdateResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) VerifyFinal(request *messages.VerifyFinalRequest) (*messages.VerifyFinalResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) VerifyRecoverInit(request *messages.VerifyRecoverInitRequest) (*messages.VerifyRecoverInitResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) VerifyRecover(request *messages.VerifyRecoverRequest) (*messages.VerifyRecoverResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
