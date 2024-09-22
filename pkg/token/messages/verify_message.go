package messages

type VerifyMessageRequest struct {
}

func (d *VerifyMessageRequest) Type() Type {
	return TypeVerifyMessageRequest
}

type WrappedVerifyMessageRequest struct {
	Type    Type
	Message VerifyMessageRequest
}

func (d *VerifyMessageRequest) Wrap() WrappedVerifyMessageRequest {
	return WrappedVerifyMessageRequest{Type: d.Type(), Message: *d}
}

type VerifyMessageResponse struct {
	ReturnValue ReturnValue
}

func (d *VerifyMessageResponse) Type() Type {
	return TypeVerifyMessageResponse
}

type WrappedVerifyMessageResponse struct {
	Type    Type
	Message VerifyMessageResponse
}

func (d *VerifyMessageResponse) Wrap() WrappedVerifyMessageResponse {
	return WrappedVerifyMessageResponse{Type: d.Type(), Message: *d}
}
