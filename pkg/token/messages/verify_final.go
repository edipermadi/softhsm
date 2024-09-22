package messages

type VerifyFinalRequest struct {
}

func (d *VerifyFinalRequest) Type() Type {
	return TypeVerifyFinalRequest
}

type WrappedVerifyFinalRequest struct {
	Type    Type
	Message VerifyFinalRequest
}

func (d *VerifyFinalRequest) Wrap() WrappedVerifyFinalRequest {
	return WrappedVerifyFinalRequest{Type: d.Type(), Message: *d}
}

type VerifyFinalResponse struct {
	ReturnValue ReturnValue
}

func (d *VerifyFinalResponse) Type() Type {
	return TypeVerifyFinalResponse
}

type WrappedVerifyFinalResponse struct {
	Type    Type
	Message VerifyFinalResponse
}

func (d *VerifyFinalResponse) Wrap() WrappedVerifyFinalResponse {
	return WrappedVerifyFinalResponse{Type: d.Type(), Message: *d}
}
