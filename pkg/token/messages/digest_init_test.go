package messages_test

import (
	"testing"

	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/stretchr/testify/require"
)

func TestDigestInitRequest_Type(t *testing.T) {
	var m messages.DigestInitRequest
	require.Equal(t, messages.TypeDigestInitRequest, m.Type())
}

func TestDigestInitResponse_Type(t *testing.T) {
	var m messages.DigestInitResponse
	require.Equal(t, messages.TypeDigestInitResponse, m.Type())
}

func TestDigestInitRequest_Wrap(t *testing.T) {
	m := messages.DigestInitRequest{
		SessionID: 1,
		Mechanism: messages.Mechanism{
			Mechanism: messages.MechanismType_SHA_1,
			Parameter: messages.MechanismParameter{
				SymmetricKey: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
				DigestLength: 2,
			},
		},
	}

	w := m.Wrap()
	require.Equal(t, m.Type(), w.Type)
	require.Equal(t, m, w.Message)
}

func TestDigestInitResponse_Wrap(t *testing.T) {
	m := messages.DigestInitResponse{
		ReturnValue: messages.ReturnValue_OK,
	}

	w := m.Wrap()
	require.Equal(t, m.Type(), w.Type)
	require.Equal(t, m, w.Message)
}
