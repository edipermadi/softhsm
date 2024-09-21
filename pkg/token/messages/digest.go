package messages

type DigestRequest struct {
	SessionID int
	Data      []byte
}

func (d *DigestRequest) Type() Type {
	return TypeDigestRequest
}

type WrappedDigestRequest struct {
	Type    Type
	Message DigestRequest
}

func (d *DigestRequest) Wrap() WrappedDigestRequest {
	return WrappedDigestRequest{Type: d.Type(), Message: *d}
}

type DigestResponse struct {
	ReturnValue ReturnValue
	Digest      []byte
}

func (d *DigestResponse) Type() Type {
	return TypeDigestResponse
}

type WrappedDigestResponse struct {
	Type    Type
	Message DigestResponse
}

func (d *DigestResponse) Wrap() WrappedDigestResponse {
	return WrappedDigestResponse{Type: d.Type(), Message: *d}
}
