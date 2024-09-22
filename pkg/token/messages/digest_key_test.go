package messages_test

import (
	"testing"

	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/stretchr/testify/require"
)

func TestDigestKeyRequest_Type(t *testing.T) {
	var m messages.DigestKeyRequest
	require.Equal(t, messages.TypeDigestKeyRequest, m.Type())
}

func TestDigestKeyResponse_Type(t *testing.T) {
	var m messages.DigestKeyResponse
	require.Equal(t, messages.TypeDigestKeyResponse, m.Type())
}

func TestDigestKeyRequest_Wrap(t *testing.T) {
	m := messages.DigestKeyRequest{
		SessionID: 1,
	}

	w := m.Wrap()
	require.Equal(t, m.Type(), w.Type)
	require.Equal(t, m, w.Message)
}

func TestDigestKeyResponse_Wrap(t *testing.T) {
	m := messages.DigestKeyResponse{
		ReturnValue: messages.ReturnValue_OK,
	}

	w := m.Wrap()
	require.Equal(t, m.Type(), w.Type)
	require.Equal(t, m, w.Message)
}
