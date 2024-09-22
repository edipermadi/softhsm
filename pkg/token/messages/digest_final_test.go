package messages_test

import (
	"testing"

	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/stretchr/testify/require"
)

func TestDigestFinalRequest_Type(t *testing.T) {
	var m messages.DigestFinalRequest
	require.Equal(t, messages.TypeDigestFinalRequest, m.Type())
}

func TestDigestFinalResponse_Type(t *testing.T) {
	var m messages.DigestFinalResponse
	require.Equal(t, messages.TypeDigestFinalResponse, m.Type())
}

func TestDigestFinalRequest_Wrap(t *testing.T) {
	m := messages.DigestFinalRequest{
		SessionID: 1,
	}

	w := m.Wrap()
	require.Equal(t, m.Type(), w.Type)
	require.Equal(t, m, w.Message)
}

func TestDigestFinalResponse_Wrap(t *testing.T) {
	m := messages.DigestFinalResponse{
		ReturnValue: messages.ReturnValue_OK,
	}

	w := m.Wrap()
	require.Equal(t, m.Type(), w.Type)
	require.Equal(t, m, w.Message)
}
