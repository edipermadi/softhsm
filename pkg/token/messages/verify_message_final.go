package messages

type VerifyMessageFinalRequest struct {
}

func (d *VerifyMessageFinalRequest) Type() Type {
	return TypeVerifyMessageFinalRequest
}

type WrappedVerifyMessageFinalRequest struct {
	Type    Type
	Message VerifyMessageFinalRequest
}

func (d *VerifyMessageFinalRequest) Wrap() WrappedVerifyMessageFinalRequest {
	return WrappedVerifyMessageFinalRequest{Type: d.Type(), Message: *d}
}

type VerifyMessageFinalResponse struct {
	ReturnValue ReturnValue
}

func (d *VerifyMessageFinalResponse) Type() Type {
	return TypeVerifyMessageFinalResponse
}

type WrappedVerifyMessageFinalResponse struct {
	Type    Type
	Message VerifyMessageFinalResponse
}

func (d *VerifyMessageFinalResponse) Wrap() WrappedVerifyMessageFinalResponse {
	return WrappedVerifyMessageFinalResponse{Type: d.Type(), Message: *d}
}
