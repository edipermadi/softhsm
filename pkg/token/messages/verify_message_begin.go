package messages

type VerifyMessageBeginRequest struct {
}

func (d *VerifyMessageBeginRequest) Type() Type {
	return TypeVerifyMessageBeginRequest
}

type WrappedVerifyMessageBeginRequest struct {
	Type    Type
	Message VerifyMessageBeginRequest
}

func (d *VerifyMessageBeginRequest) Wrap() WrappedVerifyMessageBeginRequest {
	return WrappedVerifyMessageBeginRequest{Type: d.Type(), Message: *d}
}

type VerifyMessageBeginResponse struct {
	ReturnValue ReturnValue
}

func (d *VerifyMessageBeginResponse) Type() Type {
	return TypeVerifyMessageBeginResponse
}

type WrappedVerifyMessageBeginResponse struct {
	Type    Type
	Message VerifyMessageBeginResponse
}

func (d *VerifyMessageBeginResponse) Wrap() WrappedVerifyMessageBeginResponse {
	return WrappedVerifyMessageBeginResponse{Type: d.Type(), Message: *d}
}
