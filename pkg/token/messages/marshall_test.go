package messages_test

import (
	"encoding/hex"
	"testing"

	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/stretchr/testify/require"
)

func TestMarshall(t *testing.T) {
	type testCase struct {
		Title    string
		Given    messages.Message
		Expected string
	}

	testCases := []testCase{
		{
			Title: "DigestInitRequest",
			Given: &messages.DigestInitRequest{
				SessionID: 1,
				Mechanism: messages.Mechanism{
					Mechanism: messages.MechanismType_SHA_1,
					Parameter: messages.MechanismParameter{SymmetricKey: make([]byte, 0)},
				},
			},
			Expected: "30150201693010020101300b0202022030050400020100",
		},
		{
			Title: "DigestInitResponse",
			Given: &messages.DigestInitResponse{
				ReturnValue: messages.ReturnValue_OK,
			},
			Expected: "300802016a3003020100",
		},
		{
			Title: "DigestRequest",
			Given: &messages.DigestRequest{
				SessionID: 1,
				Data:      []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			},
			Expected: "301202016b300d02010104080123456789abcdef",
		},
		{
			Title: "DigestResponse",
			Given: &messages.DigestResponse{
				ReturnValue: messages.ReturnValue_OK,
				Digest:      []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			},
			Expected: "301202016c300d02010004080123456789abcdef",
		},
		{
			Title: "DigestUpdateRequest",
			Given: &messages.DigestUpdateRequest{
				SessionID: 1,
				Data:      []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			},
			Expected: "301202016d300d02010104080123456789abcdef",
		},
		{
			Title: "DigestUpdateResponse",
			Given: &messages.DigestUpdateResponse{
				ReturnValue: messages.ReturnValue_OK,
			},
			Expected: "300802016e3003020100",
		},
		{
			Title: "DigestFinalRequest",
			Given: &messages.DigestFinalRequest{
				SessionID: 1,
			},
			Expected: "30080201713003020101",
		},
		{
			Title: "DigestFinalResponse",
			Given: &messages.DigestFinalResponse{
				ReturnValue: messages.ReturnValue_OK,
				Digest:      []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			},
			Expected: "3012020172300d02010004080123456789abcdef",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			encoded, err := messages.Marshall(tc.Given)
			require.NoError(t, err)
			require.Equal(t, tc.Expected, hex.EncodeToString(encoded))

			decoded, err := messages.Unmarshal(encoded)
			require.NoError(t, err)
			require.Equal(t, tc.Given.Type(), decoded.Type())
			switch decoded.Type() {
			case messages.TypeDigestInitRequest:
				given := tc.Given.(*messages.DigestInitRequest)
				expected, ok := decoded.(*messages.DigestInitRequest)
				require.True(t, ok)
				require.Equal(t, given, expected)
			case messages.TypeDigestInitResponse:
				given := tc.Given.(*messages.DigestInitResponse)
				expected, ok := decoded.(*messages.DigestInitResponse)
				require.True(t, ok)
				require.Equal(t, given, expected)
			case messages.TypeDigestRequest:
				given := tc.Given.(*messages.DigestRequest)
				expected, ok := decoded.(*messages.DigestRequest)
				require.True(t, ok)
				require.Equal(t, given, expected)
			case messages.TypeDigestResponse:
				given := tc.Given.(*messages.DigestResponse)
				expected, ok := decoded.(*messages.DigestResponse)
				require.True(t, ok)
				require.Equal(t, given, expected)
			case messages.TypeDigestUpdateRequest:
				given := tc.Given.(*messages.DigestUpdateRequest)
				expected, ok := decoded.(*messages.DigestUpdateRequest)
				require.True(t, ok)
				require.Equal(t, given, expected)
			case messages.TypeDigestUpdateResponse:
				given := tc.Given.(*messages.DigestUpdateResponse)
				expected, ok := decoded.(*messages.DigestUpdateResponse)
				require.True(t, ok)
				require.Equal(t, given, expected)
			case messages.TypeDigestFinalRequest:
				given := tc.Given.(*messages.DigestFinalRequest)
				expected, ok := decoded.(*messages.DigestFinalRequest)
				require.True(t, ok)
				require.Equal(t, given, expected)
			case messages.TypeDigestFinalResponse:
				given := tc.Given.(*messages.DigestFinalResponse)
				expected, ok := decoded.(*messages.DigestFinalResponse)
				require.True(t, ok)
				require.Equal(t, given, expected)
			default:
				t.Fail()
			}
		})
	}
}
