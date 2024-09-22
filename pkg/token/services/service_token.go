package services

import (
	"errors"

	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) GetSlotList(request *messages.GetSlotListRequest) (*messages.GetSlotListResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) GetSlotInfo(request *messages.GetSlotInfoRequest) (*messages.GetSlotInfoResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) GetTokenInfo(request *messages.GetTokenInfoRequest) (*messages.GetTokenInfoResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) WaitForSlotEvent(request *messages.WaitForSlotEventRequest) (*messages.WaitForSlotEventResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) GetMechanismList(request *messages.GetMechanismListRequest) (*messages.GetMechanismListResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) GetMechanismInfo(request *messages.GetMechanismInfoRequest) (*messages.GetMechanismInfoResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) InitToken(request *messages.InitTokenRequest) (*messages.InitTokenResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) InitPIN(request *messages.InitPINRequest) (*messages.InitPINResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SetPIN(request *messages.SetPINRequest) (*messages.SetPINResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
