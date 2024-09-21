package messages

type DigestFinalRequest struct {
	SessionID int
}

func (d *DigestFinalRequest) Type() Type {
	return TypeDigestFinalRequest
}

type WrappedDigestFinalRequest struct {
	Type    Type
	Message DigestFinalRequest
}

func (d *DigestFinalRequest) Wrap() WrappedDigestFinalRequest {
	return WrappedDigestFinalRequest{Type: d.Type(), Message: *d}
}

type DigestFinalResponse struct {
	ReturnValue ReturnValue
	Digest      []byte
}

func (d *DigestFinalResponse) Type() Type {
	return TypeDigestFinalResponse
}

type WrappedDigestFinalResponse struct {
	Type    Type
	Message DigestFinalResponse
}

func (d *DigestFinalResponse) Wrap() WrappedDigestFinalResponse {
	return WrappedDigestFinalResponse{Type: d.Type(), Message: *d}
}
