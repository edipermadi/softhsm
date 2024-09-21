package messages

type DigestInitRequest struct {
	SessionID int
	Mechanism Mechanism
}

func (d *DigestInitRequest) Type() Type {
	return TypeDigestInitRequest
}

type WrappedDigestInitRequest struct {
	Type    Type
	Message DigestInitRequest
}

func (d *DigestInitRequest) Wrap() WrappedDigestInitRequest {
	return WrappedDigestInitRequest{Type: d.Type(), Message: *d}
}

type DigestInitResponse struct {
	ReturnValue ReturnValue
}

func (d *DigestInitResponse) Type() Type {
	return TypeDigestInitResponse
}

type WrappedDigestInitResponse struct {
	Type    Type
	Message DigestInitResponse
}

func (d *DigestInitResponse) Wrap() WrappedDigestInitResponse {
	return WrappedDigestInitResponse{Type: d.Type(), Message: *d}
}
