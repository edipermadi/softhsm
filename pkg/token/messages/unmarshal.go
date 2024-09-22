package messages

import (
	"encoding/asn1"
	"fmt"
)

func Unmarshal(bytes []byte) (Message, error) {
	type wrapper struct {
		Type Type
	}

	var wrapped wrapper
	_, err := asn1.Unmarshal(bytes, &wrapped)
	if err != nil {
		return nil, err
	}

	switch wrapped.Type {
	case TypeDigestInitRequest:
		var decoded WrappedDigestInitRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestInitResponse:
		var decoded WrappedDigestInitResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestRequest:
		var decoded WrappedDigestRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestResponse:
		var decoded WrappedDigestResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestUpdateRequest:
		var decoded WrappedDigestUpdateRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestUpdateResponse:
		var decoded WrappedDigestUpdateResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestKeyRequest:
		var decoded WrappedDigestKeyRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestKeyResponse:
		var decoded WrappedDigestKeyResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestFinalRequest:
		var decoded WrappedDigestFinalRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestFinalResponse:
		var decoded WrappedDigestFinalResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	default:
		return nil, fmt.Errorf("unknown message type: %v", wrapped.Type)
	}
}
