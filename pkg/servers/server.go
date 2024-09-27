package servers

import (
	"sync/atomic"
	"time"

	"github.com/edipermadi/softhsm/pkg/transport/pb"
	"github.com/edipermadi/softhsm/pkg/users"
	"github.com/jellydator/ttlcache/v3"
	"github.com/sirupsen/logrus"
)

func NewServer(logger *logrus.Logger, ttl time.Duration, userService users.Service) pb.CryptographicTokenServer {
	sessionCache := ttlcache.New[uint64, Session](
		ttlcache.WithTTL[uint64, Session](ttl),
	)

	logrus.New()
	go sessionCache.Start()

	return &server{
		logger:      logger,
		ttl:         ttl,
		sessions:    sessionCache,
		userService: userService,
	}
}

type server struct {
	pb.UnimplementedCryptographicTokenServer
	logger            *logrus.Logger
	ttl               time.Duration
	userService       users.Service
	nextSessionHandle atomic.Uint64
	sessions          *ttlcache.Cache[uint64, Session]
}
