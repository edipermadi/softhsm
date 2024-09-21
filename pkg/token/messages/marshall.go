package messages

import (
	"encoding/asn1"
	"errors"
)

func Marshall(message Message) ([]byte, error) {
	switch message.Type() {
	case TypeDigestInitRequest:
		return asn1.Marshal(message.(*DigestInitRequest).Wrap())

	case TypeDigestInitResponse:
		return asn1.Marshal(message.(*DigestInitResponse).Wrap())

	case TypeDigestRequest:
		return asn1.Marshal(message.(*DigestRequest).Wrap())

	case TypeDigestResponse:
		return asn1.Marshal(message.(*DigestResponse).Wrap())

	case TypeDigestUpdateRequest:
		return asn1.Marshal(message.(*DigestUpdateRequest).Wrap())

	case TypeDigestUpdateResponse:
		return asn1.Marshal(message.(*DigestUpdateResponse).Wrap())

	case TypeDigestFinalRequest:
		return asn1.Marshal(message.(*DigestFinalRequest).Wrap())

	case TypeDigestFinalResponse:
		return asn1.Marshal(message.(*DigestFinalResponse).Wrap())

	default:
		return nil, errors.New("unexpected message type")
	}
}
