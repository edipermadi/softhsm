package messages_test

import (
	"testing"

	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/stretchr/testify/require"
)

func TestDigestRequest_Type(t *testing.T) {
	var m messages.DigestRequest
	require.Equal(t, messages.TypeDigestRequest, m.Type())
}

func TestDigestResponse_Type(t *testing.T) {
	var m messages.DigestResponse
	require.Equal(t, messages.TypeDigestResponse, m.Type())
}

func TestDigestRequest_Wrap(t *testing.T) {
	m := messages.DigestRequest{
		SessionID: 1,
		Data:      []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
	}

	w := m.Wrap()
	require.Equal(t, m.Type(), w.Type)
	require.Equal(t, m, w.Message)
}

func TestDigestResponse_Wrap(t *testing.T) {
	m := messages.DigestResponse{
		ReturnValue: messages.ReturnValue_OK,
	}

	w := m.Wrap()
	require.Equal(t, m.Type(), w.Type)
	require.Equal(t, m, w.Message)
}
