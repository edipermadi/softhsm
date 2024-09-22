package services

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) DecryptMessageInit(request *messages.DecryptMessageInitRequest) (*messages.DecryptMessageInitResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) DecryptMessage(request *messages.DecryptMessageRequest) (*messages.DecryptMessageResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) DecryptMessageBegin(request *messages.DecryptMessageBeginRequest) (*messages.DecryptMessageBeginResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) DecryptMessageNext(request *messages.DecryptMessageNextRequest) (*messages.DecryptMessageNextResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) DecryptMessageFinal(request *messages.DecryptMessageFinalRequest) (*messages.DecryptMessageFinalResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
