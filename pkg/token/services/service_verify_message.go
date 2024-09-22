package services

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) VerifyMessageInit(request *messages.VerifyMessageInitRequest) (*messages.VerifyMessageInitResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) VerifyMessage(request *messages.VerifyMessageRequest) (*messages.VerifyMessageResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) VerifyMessageBegin(request *messages.VerifyMessageBeginRequest) (*messages.VerifyMessageBeginResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) VerifyMessageNext(request *messages.VerifyMessageNextRequest) (*messages.VerifyMessageNextResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) VerifyMessageFinal(request *messages.VerifyMessageFinalRequest) (*messages.VerifyMessageFinalResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
