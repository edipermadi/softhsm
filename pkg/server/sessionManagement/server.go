package sessionManagement

import (
	"sync/atomic"
	"time"

	"github.com/edipermadi/softhsm/pkg/hsm"
	"github.com/jellydator/ttlcache/v3"
	"golang.org/x/net/context"
)

func NewServer() hsm.SessionManagementServer {
	cache := ttlcache.New[uint64, session](
		ttlcache.WithTTL[uint64, session](30 * time.Minute),
	)

	go cache.Start()

	return &server{
		cache: cache,
	}
}

type server struct {
	hsm.UnimplementedSessionManagementServer
	cache   *ttlcache.Cache[uint64, session]
	counter atomic.Uint64
}

func (s *server) OpenSession(ctx context.Context, request *hsm.OpenSessionRequest) (*hsm.OpenSessionResponse, error) {
	id := s.counter.Add(1)
	s.cache.Set(id, makeSession(id, request), 30*time.Minute)
	return &hsm.OpenSessionResponse{
		ReturnValue: hsm.ReturnValue_OK,
		Data:        &hsm.OpenSessionResponseData{SessionHandle: id}}, nil
}

func (s *server) CloseSession(ctx context.Context, request *hsm.CloseSessionRequest) (*hsm.CloseSessionResponse, error) {
	if !s.cache.Has(request.SessionHandle) {
		return &hsm.CloseSessionResponse{ReturnValue: hsm.ReturnValue_SESSION_CLOSED}, nil
	}

	s.cache.Delete(request.SessionHandle)
	return &hsm.CloseSessionResponse{ReturnValue: hsm.ReturnValue_OK}, nil
}

func (s *server) CloseAllSessions(ctx context.Context, request *hsm.CloseAllSessionsRequest) (*hsm.CloseAllSessionsResponse, error) {
	for _, k := range s.cache.Keys() {
		sess := s.cache.Get(k)
		if sess != nil && sess.Value().slotId == request.SlotId {
			s.cache.Delete(k)
		}
	}

	return &hsm.CloseAllSessionsResponse{ReturnValue: hsm.ReturnValue_OK}, nil
}

func (s *server) GetSessionInfo(ctx context.Context, request *hsm.GetSessionInfoRequest) (*hsm.GetSessionInfoResponse, error) {
	return &hsm.GetSessionInfoResponse{ReturnValue: hsm.ReturnValue_OK}, nil
}
