package messages_test

import (
	"testing"

	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomRequest_Type(t *testing.T) {
	var m messages.GenerateRandomRequest
	require.Equal(t, messages.TypeGenerateRandomRequest, m.Type())
}

func TestGenerateRandomResponse_Type(t *testing.T) {
	var m messages.GenerateRandomResponse
	require.Equal(t, messages.TypeGenerateRandomResponse, m.Type())
}

func TestGenerateRandomRequest_Wrap(t *testing.T) {
	m := messages.GenerateRandomRequest{
		SessionID: 1,
	}

	w := m.Wrap()
	require.Equal(t, m.Type(), w.Type)
	require.Equal(t, m, w.Message)
}

func TestGenerateRandomResponse_Wrap(t *testing.T) {
	m := messages.GenerateRandomResponse{
		ReturnValue: messages.ReturnValue_OK,
	}

	w := m.Wrap()
	require.Equal(t, m.Type(), w.Type)
	require.Equal(t, m, w.Message)
}
