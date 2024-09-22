package messages

type VerifyMessageNextRequest struct {
}

func (d *VerifyMessageNextRequest) Type() Type {
	return TypeVerifyMessageNextRequest
}

type WrappedVerifyMessageNextRequest struct {
	Type    Type
	Message VerifyMessageNextRequest
}

func (d *VerifyMessageNextRequest) Wrap() WrappedVerifyMessageNextRequest {
	return WrappedVerifyMessageNextRequest{Type: d.Type(), Message: *d}
}

type VerifyMessageNextResponse struct {
	ReturnValue ReturnValue
}

func (d *VerifyMessageNextResponse) Type() Type {
	return TypeVerifyMessageNextResponse
}

type WrappedVerifyMessageNextResponse struct {
	Type    Type
	Message VerifyMessageNextResponse
}

func (d *VerifyMessageNextResponse) Wrap() WrappedVerifyMessageNextResponse {
	return WrappedVerifyMessageNextResponse{Type: d.Type(), Message: *d}
}
