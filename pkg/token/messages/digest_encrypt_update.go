package messages

type DigestEncryptUpdateRequest struct {
}

func (d *DigestEncryptUpdateRequest) Type() Type {
	return TypeDigestEncryptUpdateRequest
}

type WrappedDigestEncryptUpdateRequest struct {
	Type    Type
	Message DigestEncryptUpdateRequest
}

func (d *DigestEncryptUpdateRequest) Wrap() WrappedDigestEncryptUpdateRequest {
	return WrappedDigestEncryptUpdateRequest{Type: d.Type(), Message: *d}
}

type DigestEncryptUpdateResponse struct {
	ReturnValue ReturnValue
}

func (d *DigestEncryptUpdateResponse) Type() Type {
	return TypeDigestEncryptUpdateResponse
}

type WrappedDigestEncryptUpdateResponse struct {
	Type    Type
	Message DigestEncryptUpdateResponse
}

func (d *DigestEncryptUpdateResponse) Wrap() WrappedDigestEncryptUpdateResponse {
	return WrappedDigestEncryptUpdateResponse{Type: d.Type(), Message: *d}
}
