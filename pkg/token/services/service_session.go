package services

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) OpenSession(request *messages.OpenSessionRequest) (*messages.OpenSessionResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) CloseSession(request *messages.CloseSessionRequest) (*messages.CloseSessionResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) CloseAllSessions(request *messages.CloseAllSessionsRequest) (*messages.CloseAllSessionsResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) GetSessionInfo(request *messages.GetSessionInfoRequest) (*messages.GetSessionInfoResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SessionCancel(request *messages.SessionCancelRequest) (*messages.SessionCancelResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) GetOperationState(request *messages.GetOperationStateRequest) (*messages.GetOperationStateResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SetOperationState(request *messages.SetOperationStateRequest) (*messages.SetOperationStateResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) Login(request *messages.LoginRequest) (*messages.LoginResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) LoginUser(request *messages.LoginUserRequest) (*messages.LoginUserResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) Logout(request *messages.LogoutRequest) (*messages.LogoutResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
