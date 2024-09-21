package messages

type DigestUpdateRequest struct {
	SessionID int
	Data      []byte
}

func (d *DigestUpdateRequest) Type() Type {
	return TypeDigestUpdateRequest
}

type WrappedDigestUpdateRequest struct {
	Type    Type
	Message DigestUpdateRequest
}

func (d *DigestUpdateRequest) Wrap() WrappedDigestUpdateRequest {
	return WrappedDigestUpdateRequest{Type: d.Type(), Message: *d}
}

type DigestUpdateResponse struct {
	ReturnValue ReturnValue
}

func (d *DigestUpdateResponse) Type() Type {
	return TypeDigestUpdateResponse
}

type WrappedDigestUpdateResponse struct {
	Type    Type
	Message DigestUpdateResponse
}

func (d *DigestUpdateResponse) Wrap() WrappedDigestUpdateResponse {
	return WrappedDigestUpdateResponse{Type: d.Type(), Message: *d}
}
