package messages

type DigestKeyRequest struct {
	SessionID int
	KeyID     int
}

func (d *DigestKeyRequest) Type() Type {
	return TypeDigestKeyRequest
}

type WrappedDigestKeyRequest struct {
	Type    Type
	Message DigestKeyRequest
}

func (d *DigestKeyRequest) Wrap() WrappedDigestKeyRequest {
	return WrappedDigestKeyRequest{Type: d.Type(), Message: *d}
}

type DigestKeyResponse struct {
	ReturnValue ReturnValue
	Digest      []byte
}

func (d *DigestKeyResponse) Type() Type {
	return TypeDigestKeyResponse
}

type WrappedDigestKeyResponse struct {
	Type    Type
	Message DigestKeyResponse
}

func (d *DigestKeyResponse) Wrap() WrappedDigestKeyResponse {
	return WrappedDigestKeyResponse{Type: d.Type(), Message: *d}
}
