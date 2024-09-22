package services

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) GetFunctionStatus(request *messages.GetFunctionStatusRequest) (*messages.GetFunctionStatusResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) CancelFunction(request *messages.CancelFunctionRequest) (*messages.CancelFunctionResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
