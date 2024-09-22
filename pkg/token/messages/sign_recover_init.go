package messages

type SignRecoverInitRequest struct {
}

func (d *SignRecoverInitRequest) Type() Type {
	return TypeSignRecoverInitRequest
}

type WrappedSignRecoverInitRequest struct {
	Type    Type
	Message SignRecoverInitRequest
}

func (d *SignRecoverInitRequest) Wrap() WrappedSignRecoverInitRequest {
	return WrappedSignRecoverInitRequest{Type: d.Type(), Message: *d}
}

type SignRecoverInitResponse struct {
	ReturnValue ReturnValue
}

func (d *SignRecoverInitResponse) Type() Type {
	return TypeSignRecoverInitResponse
}

type WrappedSignRecoverInitResponse struct {
	Type    Type
	Message SignRecoverInitResponse
}

func (d *SignRecoverInitResponse) Wrap() WrappedSignRecoverInitResponse {
	return WrappedSignRecoverInitResponse{Type: d.Type(), Message: *d}
}
