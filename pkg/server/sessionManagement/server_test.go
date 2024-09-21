package sessionManagement_test

import (
	"context"
	"encoding/asn1"
	"encoding/hex"
	"log"
	"net"
	"testing"

	"github.com/edipermadi/softhsm/pkg/hsm"
	"github.com/edipermadi/softhsm/pkg/server/sessionManagement"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (hsm.SessionManagementClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	hsm.RegisterSessionManagementServer(baseServer, sessionManagement.NewServer())
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

	client := hsm.NewSessionManagementClient(conn)

	return client, closer
}

func TestServer_OpenSession(t *testing.T) {
	ctx := context.TODO()
	client, closer := server(ctx)
	defer closer()

	for i := uint64(1); i <= uint64(10); i++ {
		response, err := client.OpenSession(ctx, &hsm.OpenSessionRequest{
			SlotId: 1,
			Flags:  2,
		})
		require.NoError(t, err)
		require.Equal(t, hsm.ReturnValue_OK, response.GetReturnValue())
		require.Equal(t, i, response.Data.SessionHandle)
	}
}

func TestServer_CloseSession(t *testing.T) {
	ctx := context.TODO()
	client, closer := server(ctx)
	defer closer()

	{
		response, err := client.OpenSession(ctx, &hsm.OpenSessionRequest{
			SlotId: 1,
			Flags:  2,
		})
		require.NoError(t, err)
		require.Equal(t, hsm.ReturnValue_OK, response.GetReturnValue())
		require.Equal(t, uint64(1), response.Data.SessionHandle)
	}

	{
		response, err := client.CloseSession(ctx, &hsm.CloseSessionRequest{SessionHandle: uint64(1)})
		require.NoError(t, err)
		require.Equal(t, hsm.ReturnValue_OK, response.GetReturnValue())
	}
}

func TestServer_CloseAllSessions(t *testing.T) {
	ctx := context.TODO()
	client, closer := server(ctx)
	defer closer()

	for i := uint64(1); i <= uint64(10); i++ {
		response, err := client.OpenSession(ctx, &hsm.OpenSessionRequest{
			SlotId: 1,
			Flags:  2,
		})
		require.NoError(t, err)
		require.Equal(t, hsm.ReturnValue_OK, response.GetReturnValue())
		require.Equal(t, i, response.Data.SessionHandle)
	}

	{
		response, err := client.CloseAllSessions(ctx, &hsm.CloseAllSessionsRequest{SlotId: uint64(1)})
		require.NoError(t, err)
		require.Equal(t, hsm.ReturnValue_OK, response.GetReturnValue())
	}
}

func TestEncodeASN1(t *testing.T) {
	type Data struct {
		ID   int64
		Name string
	}

	data := Data{1, "edi"}
	marshalled, err := asn1.Marshal(data)
	require.NoError(t, err)
	t.Logf("result = %s", hex.EncodeToString(marshalled))
}
