package servers

import (
	"context"
	"hash"

	"github.com/edipermadi/softhsm/pkg/transport/pb"
)

type Session struct {
	hash hash.Hash
}

func (s *server) OpenSession(ctx context.Context, request *pb.OpenSessionRequest) (*pb.OpenSessionResponse, error) {
	sessionHandle := s.nextSessionHandle.Add(1)
	s.sessions.Set(sessionHandle, Session{}, s.ttl)
	return &pb.OpenSessionResponse{
		ReturnValue:   pb.ReturnValue_OK,
		SessionHandle: sessionHandle,
	}, nil
}

func (s *server) CloseSession(ctx context.Context, request *pb.CloseSessionRequest) (*pb.CloseSessionResponse, error) {
	if request.GetSessionHandle() < 1 {
		return &pb.CloseSessionResponse{ReturnValue: pb.ReturnValue_SESSION_HANDLE_INVALID}, nil
	}

	if !s.sessions.Has(request.GetSessionHandle()) {
		return &pb.CloseSessionResponse{ReturnValue: pb.ReturnValue_SESSION_CLOSED}, nil
	}

	s.sessions.Delete(request.GetSessionHandle())
	return &pb.CloseSessionResponse{ReturnValue: pb.ReturnValue_OK}, nil
}
