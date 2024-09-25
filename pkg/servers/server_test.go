package servers_test

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"time"

	"github.com/edipermadi/softhsm/pkg/servers"
	"github.com/edipermadi/softhsm/pkg/transport/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func testServer(ctx context.Context) (pb.CryptographicTokenClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	logger := logrus.New()
	baseServer := grpc.NewServer()
	pb.RegisterCryptographicTokenServer(baseServer, servers.NewServer(logger, 30*time.Minute))
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

	return client, closer
}
