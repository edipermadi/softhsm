package services

import (
	"errors"

	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) Initialize(request *messages.InitializeRequest) (*messages.InitializeResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) Finalize(request *messages.FinalizeRequest) (*messages.FinalizeResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) GetInfo(request *messages.GetInfoRequest) (*messages.GetInfoResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) GetFunctionList(request *messages.GetFunctionListRequest) (*messages.GetFunctionListResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) GetInterfaceList(request *messages.GetInterfaceListRequest) (*messages.GetInterfaceListResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) GetInterface(request *messages.GetInterfaceRequest) (*messages.GetInterfaceResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
