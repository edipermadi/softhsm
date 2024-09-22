package services

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/edipermadi/softhsm/pkg/token/sessions"
)

func (s *service) OpenSession(request *messages.OpenSessionRequest) (*messages.OpenSessionResponse, error) {
	sessionID := int(s.sessionSequence.Add(1))
	s.sessions[sessionID] = sessions.Session{Flags: request.Flags}
	return &messages.OpenSessionResponse{
		ReturnValue: messages.ReturnValue_OK,
		SessionID:   sessionID,
	}, nil
}

func (s *service) CloseSession(request *messages.CloseSessionRequest) (*messages.CloseSessionResponse, error) {
	if _, found := s.sessions[request.SessionID]; !found {
		return &messages.CloseSessionResponse{ReturnValue: messages.ReturnValue_SESSION_CLOSED}, nil
	}

	delete(s.sessions, request.SessionID)

	return &messages.CloseSessionResponse{
		ReturnValue: messages.ReturnValue_OK,
	}, nil
}

func (s *service) CloseAllSessions(request *messages.CloseAllSessionsRequest) (*messages.CloseAllSessionsResponse, error) {
	for sessionID := range s.sessions {
		delete(s.sessions, sessionID)
	}

	return &messages.CloseAllSessionsResponse{
		ReturnValue: messages.ReturnValue_OK,
	}, nil
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
