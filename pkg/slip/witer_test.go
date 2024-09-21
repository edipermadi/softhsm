package slip_test

import (
	"bytes"
	"testing"

	"github.com/edipermadi/softhsm/pkg/slip"
	"github.com/stretchr/testify/require"
)

func TestWriter_Write(t *testing.T) {
	type testCase struct {
		Title    string
		Given    []byte
		Expected []byte
	}

	testCases := []testCase{
		{
			Title:    "Esc",
			Given:    []byte{slip.Esc},
			Expected: []byte{slip.End, slip.Esc, slip.EscEsc, slip.End},
		},
		{
			Title:    "End",
			Given:    []byte{slip.End},
			Expected: []byte{slip.End, slip.Esc, slip.EscEnd, slip.End},
		},
		{
			Title:    "Esc_End",
			Given:    []byte{slip.Esc, slip.End},
			Expected: []byte{slip.End, slip.Esc, slip.EscEsc, slip.Esc, slip.EscEnd, slip.End},
		},
		{
			Title:    "End_Esc",
			Given:    []byte{slip.End, slip.Esc},
			Expected: []byte{slip.End, slip.Esc, slip.EscEnd, slip.Esc, slip.EscEsc, slip.End},
		},
		{
			Title:    "EscEsc",
			Given:    []byte{slip.EscEsc},
			Expected: []byte{slip.End, slip.EscEsc, slip.End},
		},
		{
			Title:    "EscEnd",
			Given:    []byte{slip.EscEnd},
			Expected: []byte{slip.End, slip.EscEnd, slip.End},
		},
		{
			Title:    "EscEsc_EscEnd",
			Given:    []byte{slip.EscEsc, slip.EscEnd},
			Expected: []byte{slip.End, slip.EscEsc, slip.EscEnd, slip.End},
		},
		{
			Title:    "EscEnd_EscEsc",
			Given:    []byte{slip.EscEnd, slip.EscEsc},
			Expected: []byte{slip.End, slip.EscEnd, slip.EscEsc, slip.End},
		},
		{
			Title:    "NoEscape",
			Given:    []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			Expected: []byte{slip.End, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, slip.End},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			var encoded bytes.Buffer
			writer := slip.NewWriter(&encoded)
			written, err := writer.Write(tc.Given)
			require.NoError(t, err)
			require.Equal(t, len(tc.Expected), written)
			require.Equal(t, tc.Expected, encoded.Bytes())
		})
	}
}
