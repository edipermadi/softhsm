package slip_test

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/edipermadi/softhsm/pkg/slip"
	"github.com/stretchr/testify/require"
)

func TestReader_Read_Random(t *testing.T) {
	data := make([]byte, 1024)
	length, err := rand.Read(data)
	require.NoError(t, err)
	require.Equal(t, len(data), length)

	var encoded bytes.Buffer
	writer := slip.NewWriter(&encoded)
	written, err := writer.Write(data)
	require.NoError(t, err)

	decoded := make([]byte, len(data))
	decoder := slip.NewReader(bytes.NewReader(encoded.Bytes()))
	length, err = decoder.Read(decoded)
	require.NoError(t, err)
	require.Equal(t, length, written)
}

func TestReader_Read(t *testing.T) {
	type testCase struct {
		Title    string
		Given    []byte
		Expected []byte
	}

	testCases := []testCase{
		{
			Title:    "Esc",
			Given:    []byte{slip.End, slip.Esc, slip.EscEsc, slip.End},
			Expected: []byte{slip.Esc},
		},
		{
			Title:    "End",
			Given:    []byte{slip.End, slip.Esc, slip.EscEnd, slip.End},
			Expected: []byte{slip.End},
		},
		{
			Title:    "Esc_End",
			Given:    []byte{slip.End, slip.Esc, slip.EscEsc, slip.Esc, slip.EscEnd, slip.End},
			Expected: []byte{slip.Esc, slip.End},
		},
		{
			Title:    "End_Esc",
			Given:    []byte{slip.End, slip.Esc, slip.EscEnd, slip.Esc, slip.EscEsc, slip.End},
			Expected: []byte{slip.End, slip.Esc},
		},
		{
			Title:    "EscEsc",
			Given:    []byte{slip.End, slip.EscEsc, slip.End},
			Expected: []byte{slip.EscEsc},
		},
		{
			Title:    "EscEnd",
			Given:    []byte{slip.End, slip.EscEnd, slip.End},
			Expected: []byte{slip.EscEnd},
		},
		{
			Title:    "EscEsc_EscEnd",
			Given:    []byte{slip.End, slip.EscEsc, slip.EscEnd, slip.End},
			Expected: []byte{slip.EscEsc, slip.EscEnd},
		},
		{
			Title:    "EscEnd_EscEsc",
			Given:    []byte{slip.End, slip.EscEnd, slip.EscEsc, slip.End},
			Expected: []byte{slip.EscEnd, slip.EscEsc},
		},
		{
			Title:    "NoEscape",
			Given:    []byte{slip.End, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, slip.End},
			Expected: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			decoded := make([]byte, len(tc.Expected))
			reader := slip.NewReader(bytes.NewReader(tc.Given))
			read, err := reader.Read(decoded)
			require.NoError(t, err)
			require.Equal(t, len(tc.Given), read)
			require.Equal(t, tc.Expected, decoded)
		})
	}
}
