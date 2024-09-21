package messages_test

import (
	"testing"

	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/stretchr/testify/require"
)

func TestDigestUpdateRequest_Type(t *testing.T) {
	var m messages.DigestUpdateRequest
	require.Equal(t, messages.TypeDigestUpdateRequest, m.Type())
}

func TestDigestUpdateResponse_Type(t *testing.T) {
	var m messages.DigestUpdateResponse
	require.Equal(t, messages.TypeDigestUpdateResponse, m.Type())
}

func TestDigestUpdateRequest_Wrap(t *testing.T) {
	m := messages.DigestUpdateRequest{
		SessionID: 1,
		Data:      []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
	}

	w := m.Wrap()
	require.Equal(t, m.Type(), w.Type)
	require.Equal(t, m, w.Message)
}

func TestDigestUpdateResponse_Wrap(t *testing.T) {
	m := messages.DigestUpdateResponse{
		ReturnValue: messages.ReturnValue_OK,
	}

	w := m.Wrap()
	require.Equal(t, m.Type(), w.Type)
	require.Equal(t, m, w.Message)
}
