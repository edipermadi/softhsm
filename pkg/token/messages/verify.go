package messages

type VerifyRequest struct {
}

func (d *VerifyRequest) Type() Type {
	return TypeVerifyRequest
}

type WrappedVerifyRequest struct {
	Type    Type
	Message VerifyRequest
}

func (d *VerifyRequest) Wrap() WrappedVerifyRequest {
	return WrappedVerifyRequest{Type: d.Type(), Message: *d}
}

type VerifyResponse struct {
	ReturnValue ReturnValue
}

func (d *VerifyResponse) Type() Type {
	return TypeVerifyResponse
}

type WrappedVerifyResponse struct {
	Type    Type
	Message VerifyResponse
}

func (d *VerifyResponse) Wrap() WrappedVerifyResponse {
	return WrappedVerifyResponse{Type: d.Type(), Message: *d}
}
