package services

import (
	"errors"
	"slices"

	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/edipermadi/softhsm/pkg/token/sessions"
	"github.com/edipermadi/softhsm/pkg/token/users"
)

func (s *service) OpenSession(request *messages.OpenSessionRequest) (*messages.OpenSessionResponse, error) {
	if request == nil {
		return &messages.OpenSessionResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	s.sessionsMutex.Lock()
	defer s.sessionsMutex.Unlock()

	sessionID := int(s.sessionSequence.Add(1))
	s.sessions[sessionID] = sessions.Session{Flags: request.Flags}
	return &messages.OpenSessionResponse{
		ReturnValue: messages.ReturnValue_OK,
		SessionID:   sessionID,
	}, nil
}

func (s *service) CloseSession(request *messages.CloseSessionRequest) (*messages.CloseSessionResponse, error) {
	if request == nil {
		return &messages.CloseSessionResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	if request.SessionID < 1 {
		return &messages.CloseSessionResponse{ReturnValue: messages.ReturnValue_SESSION_HANDLE_INVALID}, nil
	}

	s.sessionsMutex.Lock()
	defer s.sessionsMutex.Unlock()

	if _, found := s.sessions[request.SessionID]; !found {
		return &messages.CloseSessionResponse{ReturnValue: messages.ReturnValue_SESSION_CLOSED}, nil
	}

	delete(s.sessions, request.SessionID)

	return &messages.CloseSessionResponse{
		ReturnValue: messages.ReturnValue_OK,
	}, nil
}

func (s *service) CloseAllSessions(request *messages.CloseAllSessionsRequest) (*messages.CloseAllSessionsResponse, error) {
	if request == nil {
		return &messages.CloseAllSessionsResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	s.sessionsMutex.Lock()
	defer s.sessionsMutex.Unlock()

	for sessionID := range s.sessions {
		delete(s.sessions, sessionID)
	}

	return &messages.CloseAllSessionsResponse{
		ReturnValue: messages.ReturnValue_OK,
	}, nil
}

func (s *service) GetSessionInfo(request *messages.GetSessionInfoRequest) (*messages.GetSessionInfoResponse, error) {
	if request == nil {
		return &messages.GetSessionInfoResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	s.sessionsMutex.Lock()
	defer s.sessionsMutex.Unlock()

	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SessionCancel(request *messages.SessionCancelRequest) (*messages.SessionCancelResponse, error) {
	if request == nil {
		return &messages.SessionCancelResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) GetOperationState(request *messages.GetOperationStateRequest) (*messages.GetOperationStateResponse, error) {
	if request == nil {
		return &messages.GetOperationStateResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SetOperationState(request *messages.SetOperationStateRequest) (*messages.SetOperationStateResponse, error) {
	if request == nil {
		return &messages.SetOperationStateResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) Login(request *messages.LoginRequest) (*messages.LoginResponse, error) {
	if request == nil {
		return &messages.LoginResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	if request.SessionID < 1 {
		return &messages.LoginResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	acceptableUserTypes := []users.Type{
		users.TypeSecurityOfficer,
		users.TypeUser,
	}

	if !slices.Contains(acceptableUserTypes, request.UserType) {
		return &messages.LoginResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	if len(request.Pin) < 1 {
		return &messages.LoginResponse{ReturnValue: messages.ReturnValue_PIN_INVALID}, nil
	}

	s.sessionsMutex.Lock()
	defer s.sessionsMutex.Unlock()

	session, found := s.sessions[request.SessionID]
	if !found {
		return &messages.LoginResponse{ReturnValue: messages.ReturnValue_SESSION_CLOSED}, nil
	}

	if session.UserType != users.TypeInvalid {
		return &messages.LoginResponse{ReturnValue: messages.ReturnValue_USER_ALREADY_LOGGED_IN}, nil
	}

	if err := s.userService.Authenticate("default", request.Pin); err != nil {
		return &messages.LoginResponse{ReturnValue: messages.ReturnValue_PIN_INCORRECT}, nil
	}

	session.UserType = request.UserType
	session.Username = "default"
	s.sessions[request.SessionID] = session
	return &messages.LoginResponse{ReturnValue: messages.ReturnValue_OK}, nil
}

func (s *service) LoginUser(request *messages.LoginUserRequest) (*messages.LoginUserResponse, error) {
	if request == nil {
		return &messages.LoginUserResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	if request.SessionID < 1 {
		return &messages.LoginUserResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	if (request.UserType != users.TypeSecurityOfficer) && (request.UserType != users.TypeUser) {
		return &messages.LoginUserResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	if len(request.Pin) < 1 {
		return &messages.LoginUserResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	if len(request.Username) < 1 {
		return &messages.LoginUserResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	s.sessionsMutex.Lock()
	defer s.sessionsMutex.Unlock()

	session, found := s.sessions[request.SessionID]
	if !found {
		return &messages.LoginUserResponse{ReturnValue: messages.ReturnValue_SESSION_CLOSED}, nil
	}

	if session.UserType != users.TypeInvalid {
		return &messages.LoginUserResponse{ReturnValue: messages.ReturnValue_USER_ALREADY_LOGGED_IN}, nil
	}

	if err := s.userService.Authenticate(request.Username, request.Pin); err != nil {
		return &messages.LoginUserResponse{ReturnValue: messages.ReturnValue_PIN_INCORRECT}, nil
	}

	session.UserType = request.UserType
	session.Username = request.Username
	s.sessions[request.SessionID] = session
	return &messages.LoginUserResponse{ReturnValue: messages.ReturnValue_OK}, nil
}

func (s *service) Logout(request *messages.LogoutRequest) (*messages.LogoutResponse, error) {
	if request == nil {
		return &messages.LogoutResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	if request.SessionID < 1 {
		return &messages.LogoutResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	s.sessionsMutex.Lock()
	defer s.sessionsMutex.Unlock()

	_, found := s.sessions[request.SessionID]
	if !found {
		return &messages.LogoutResponse{ReturnValue: messages.ReturnValue_SESSION_CLOSED}, nil
	}

	s.sessions[request.SessionID] = sessions.Session{}

	// TODO add implementation
	return nil, errors.New("not implemented")
}
