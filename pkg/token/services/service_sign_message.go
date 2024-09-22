package services

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) SignMessageInit(request *messages.SignMessageInitRequest) (*messages.SignMessageInitResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SignMessage(request *messages.SignMessageRequest) (*messages.SignMessageResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SignMessageBegin(request *messages.SignMessageBeginRequest) (*messages.SignMessageBeginResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SignMessageNext(request *messages.SignMessageNextRequest) (*messages.SignMessageNextResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SignMessageFinal(request *messages.SignMessageFinalRequest) (*messages.SignMessageFinalResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
