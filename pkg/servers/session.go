package servers

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"hash"

	"github.com/edipermadi/softhsm/pkg/transport/pb"
	"github.com/edipermadi/softhsm/pkg/users"
)

type Session struct {
	hash       hash.Hash
	isLoggedIn bool
}

func (s *server) OpenSession(ctx context.Context, request *pb.OpenSessionRequest) (*pb.OpenSessionResponse, error) {
	logger := s.logger.WithContext(ctx).WithFields(logrus.Fields{
		"action": "OpenSession",
	})

	logger.Debug("opening session")

	sessionHandle := s.nextSessionHandle.Add(1)
	s.sessions.Set(sessionHandle, Session{}, s.ttl)
	return &pb.OpenSessionResponse{
		ReturnValue:   pb.ReturnValue_OK,
		SessionHandle: sessionHandle,
	}, nil
}

func (s *server) CloseSession(ctx context.Context, request *pb.CloseSessionRequest) (*pb.CloseSessionResponse, error) {
	logger := s.logger.WithContext(ctx).WithFields(logrus.Fields{
		"action":         "CloseSession",
		"session_handle": request.SessionHandle,
	})

	logger.Debug("closing session")

	if request.GetSessionHandle() < 1 {
		return &pb.CloseSessionResponse{ReturnValue: pb.ReturnValue_SESSION_HANDLE_INVALID}, nil
	}

	if !s.sessions.Has(request.GetSessionHandle()) {
		return &pb.CloseSessionResponse{ReturnValue: pb.ReturnValue_SESSION_CLOSED}, nil
	}

	s.sessions.Delete(request.GetSessionHandle())
	return &pb.CloseSessionResponse{ReturnValue: pb.ReturnValue_OK}, nil
}

func (s *server) CloseAllSessions(ctx context.Context, request *pb.CloseAllSessionsRequest) (*pb.CloseAllSessionsResponse, error) {
	logger := s.logger.WithContext(ctx).WithFields(logrus.Fields{
		"action": "CloseAllSessions",
	})

	logger.Debug("closing all sessions")

	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *server) GetSessionInfo(ctx context.Context, request *pb.GetSessionInfoRequest) (*pb.GetSessionInfoResponse, error) {
	logger := s.logger.WithContext(ctx).WithFields(logrus.Fields{
		"action":         "GetSessionInfo",
		"session_handle": request.SessionHandle,
	})

	logger.Debug("getting session info")

	// TODO add implementation
	return nil, errors.New("not implemented")

}
func (s *server) SessionCancel(ctx context.Context, request *pb.SessionCancelRequest) (*pb.SessionCancelResponse, error) {
	logger := s.logger.WithContext(ctx).WithFields(logrus.Fields{
		"action":         "SessionCancel",
		"session_handle": request.SessionHandle,
	})

	logger.Debug("cancelling session info")

	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *server) GetOperationState(ctx context.Context, request *pb.GetOperationStateRequest) (*pb.GetOperationStateResponse, error) {
	logger := s.logger.WithContext(ctx).WithFields(logrus.Fields{
		"action":         "GetOperationState",
		"session_handle": request.SessionHandle,
	})

	logger.Debug("getting operation state")

	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *server) SetOperationState(ctx context.Context, request *pb.SetOperationStateRequest) (*pb.SetOperationStateResponse, error) {
	logger := s.logger.WithContext(ctx).WithFields(logrus.Fields{
		"action":         "SetOperationState",
		"session_handle": request.SessionHandle,
	})

	logger.Debug("setting operation state")

	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *server) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	logger := s.logger.WithContext(ctx).WithFields(logrus.Fields{
		"action":         "Login",
		"session_handle": request.SessionHandle,
	})

	logger.Debug("logging in")

	if request.GetSessionHandle() < 1 {
		return &pb.LoginResponse{ReturnValue: pb.ReturnValue_SESSION_HANDLE_INVALID}, nil
	}

	if len(request.GetPin()) < 1 {
		return &pb.LoginResponse{ReturnValue: pb.ReturnValue_ARGUMENTS_BAD}, nil
	}

	if !s.sessions.Has(request.GetSessionHandle()) {
		return &pb.LoginResponse{ReturnValue: pb.ReturnValue_SESSION_CLOSED}, nil
	}

	session := s.sessions.Get(request.GetSessionHandle()).Value()

	userTypes := map[pb.UserType]users.Type{
		pb.UserType_SO:   users.TypeSecurityOfficer,
		pb.UserType_USER: users.TypeUser,
	}

	userNames := map[pb.UserType]string{
		pb.UserType_SO:   users.DefaultSecurityOfficer,
		pb.UserType_USER: users.DefaultUser,
	}

	userType, found := userTypes[request.UserType]
	if !found {
		userType = users.TypeInvalid
	}

	userName, found := userNames[request.UserType]
	if !found {
		userName = users.DefaultSecurityOfficer
	}

	err := s.userService.Authenticate(ctx, userType, userName, request.Pin)
	switch {
	case errors.Is(err, users.ErrPasswordIsIncorrect):
		return &pb.LoginResponse{ReturnValue: pb.ReturnValue_PIN_INCORRECT}, nil
	case errors.Is(err, users.ErrPasswordIsExpired):
		return &pb.LoginResponse{ReturnValue: pb.ReturnValue_PIN_EXPIRED}, nil
	case err != nil:
		logger.WithError(err).Error("failed to authenticate user")
		return &pb.LoginResponse{ReturnValue: pb.ReturnValue_GENERAL_ERROR}, nil
	}

	session.isLoggedIn = true
	s.sessions.Set(request.SessionHandle, session, s.ttl)
	return &pb.LoginResponse{ReturnValue: pb.ReturnValue_OK}, nil
}

func (s *server) LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	logger := s.logger.WithContext(ctx).WithFields(logrus.Fields{
		"action":         "LoginUser",
		"session_handle": request.SessionHandle,
	})

	logger.Debug("logging in")

	if request.GetSessionHandle() < 1 {
		return &pb.LoginUserResponse{ReturnValue: pb.ReturnValue_SESSION_HANDLE_INVALID}, nil
	}

	if len(request.GetUsername()) < 1 {
		return &pb.LoginUserResponse{ReturnValue: pb.ReturnValue_ARGUMENTS_BAD}, nil
	}

	if len(request.GetPin()) < 1 {
		return &pb.LoginUserResponse{ReturnValue: pb.ReturnValue_ARGUMENTS_BAD}, nil
	}

	if !s.sessions.Has(request.GetSessionHandle()) {
		return &pb.LoginUserResponse{ReturnValue: pb.ReturnValue_SESSION_CLOSED}, nil
	}

	session := s.sessions.Get(request.GetSessionHandle()).Value()

	userTypes := map[pb.UserType]users.Type{
		pb.UserType_SO:   users.TypeSecurityOfficer,
		pb.UserType_USER: users.TypeUser,
	}

	userType, found := userTypes[request.UserType]
	if !found {
		userType = users.TypeInvalid
	}

	err := s.userService.Authenticate(ctx, userType, request.Username, request.Pin)
	switch {
	case errors.Is(err, users.ErrPasswordIsIncorrect):
		return &pb.LoginUserResponse{ReturnValue: pb.ReturnValue_PIN_INCORRECT}, nil
	case errors.Is(err, users.ErrPasswordIsExpired):
		return &pb.LoginUserResponse{ReturnValue: pb.ReturnValue_PIN_EXPIRED}, nil
	case err != nil:
		logger.WithError(err).Error("failed to authenticate user")
		return &pb.LoginUserResponse{ReturnValue: pb.ReturnValue_GENERAL_ERROR}, nil
	}

	session.isLoggedIn = true
	s.sessions.Set(request.SessionHandle, session, s.ttl)
	return &pb.LoginUserResponse{ReturnValue: pb.ReturnValue_OK}, nil
}

func (s *server) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
