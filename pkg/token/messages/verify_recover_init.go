package messages

type VerifyRecoverInitRequest struct {
}

func (d *VerifyRecoverInitRequest) Type() Type {
	return TypeVerifyRecoverInitRequest
}

type WrappedVerifyRecoverInitRequest struct {
	Type    Type
	Message VerifyRecoverInitRequest
}

func (d *VerifyRecoverInitRequest) Wrap() WrappedVerifyRecoverInitRequest {
	return WrappedVerifyRecoverInitRequest{Type: d.Type(), Message: *d}
}

type VerifyRecoverInitResponse struct {
	ReturnValue ReturnValue
}

func (d *VerifyRecoverInitResponse) Type() Type {
	return TypeVerifyRecoverInitResponse
}

type WrappedVerifyRecoverInitResponse struct {
	Type    Type
	Message VerifyRecoverInitResponse
}

func (d *VerifyRecoverInitResponse) Wrap() WrappedVerifyRecoverInitResponse {
	return WrappedVerifyRecoverInitResponse{Type: d.Type(), Message: *d}
}
