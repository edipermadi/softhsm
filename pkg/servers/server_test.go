package servers_test

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/edipermadi/softhsm/pkg/servers"
	"github.com/edipermadi/softhsm/pkg/transport/pb"
	"github.com/edipermadi/softhsm/pkg/users"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type mockedUserService struct {
	mock.Mock
}

func (m *mockedUserService) Authenticate(ctx context.Context, userType users.Type, username string, password string) error {
	args := m.Called(ctx, userType, username, password)
	return args.Error(0)
}

func (m *mockedUserService) RecordLoginAttempt(ctx context.Context, userId int64, succeeded bool) error {
	args := m.Called(ctx, userId, succeeded)
	return args.Error(0)
}

func testServer(ctx context.Context) (*mockedUserService, pb.CryptographicTokenClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	userService := &mockedUserService{}
	logger := logrus.New()
	baseServer := grpc.NewServer()
	pb.RegisterCryptographicTokenServer(baseServer, servers.NewServer(logger, 30*time.Minute, userService))
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewCryptographicTokenClient(conn)

	return userService, client, closer
}
