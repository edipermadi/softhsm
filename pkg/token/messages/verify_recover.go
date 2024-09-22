package messages

type VerifyRecoverRequest struct {
}

func (d *VerifyRecoverRequest) Type() Type {
	return TypeVerifyRecoverRequest
}

type WrappedVerifyRecoverRequest struct {
	Type    Type
	Message VerifyRecoverRequest
}

func (d *VerifyRecoverRequest) Wrap() WrappedVerifyRecoverRequest {
	return WrappedVerifyRecoverRequest{Type: d.Type(), Message: *d}
}

type VerifyRecoverResponse struct {
	ReturnValue ReturnValue
}

func (d *VerifyRecoverResponse) Type() Type {
	return TypeVerifyRecoverResponse
}

type WrappedVerifyRecoverResponse struct {
	Type    Type
	Message VerifyRecoverResponse
}

func (d *VerifyRecoverResponse) Wrap() WrappedVerifyRecoverResponse {
	return WrappedVerifyRecoverResponse{Type: d.Type(), Message: *d}
}
