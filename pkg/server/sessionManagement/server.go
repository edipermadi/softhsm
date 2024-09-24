package sessionManagement

import (
	"sync/atomic"
	"time"

	"github.com/edipermadi/softhsm/pkg/transport/pb"
	"github.com/jellydator/ttlcache/v3"
	"golang.org/x/net/context"
)

func NewServer() pb.SessionManagementServer {
	cache := ttlcache.New[uint64, session](
		ttlcache.WithTTL[uint64, session](30 * time.Minute),
	)

	go cache.Start()

	return &server{
		cache: cache,
	}
}

type server struct {
	pb.UnimplementedSessionManagementServer
	cache   *ttlcache.Cache[uint64, session]
	counter atomic.Uint64
}

func (s *server) OpenSession(ctx context.Context, request *pb.OpenSessionRequest) (*pb.OpenSessionResponse, error) {
	id := s.counter.Add(1)
	s.cache.Set(id, makeSession(id, request), 30*time.Minute)
	return &pb.OpenSessionResponse{
		ReturnValue: pb.ReturnValue_OK,
		Data:        &pb.OpenSessionResponseData{SessionHandle: id}}, nil
}

func (s *server) CloseSession(ctx context.Context, request *pb.CloseSessionRequest) (*pb.CloseSessionResponse, error) {
	if !s.cache.Has(request.SessionHandle) {
		return &pb.CloseSessionResponse{ReturnValue: pb.ReturnValue_SESSION_CLOSED}, nil
	}

	s.cache.Delete(request.SessionHandle)
	return &pb.CloseSessionResponse{ReturnValue: pb.ReturnValue_OK}, nil
}

func (s *server) CloseAllSessions(ctx context.Context, request *pb.CloseAllSessionsRequest) (*pb.CloseAllSessionsResponse, error) {
	for _, k := range s.cache.Keys() {
		sess := s.cache.Get(k)
		if sess != nil && sess.Value().slotId == request.SlotId {
			s.cache.Delete(k)
		}
	}

	return &pb.CloseAllSessionsResponse{ReturnValue: pb.ReturnValue_OK}, nil
}

func (s *server) GetSessionInfo(ctx context.Context, request *pb.GetSessionInfoRequest) (*pb.GetSessionInfoResponse, error) {
	return &pb.GetSessionInfoResponse{ReturnValue: pb.ReturnValue_OK}, nil
}
