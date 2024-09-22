package messages

type SignRecoverRequest struct {
}

func (d *SignRecoverRequest) Type() Type {
	return TypeSignRecoverRequest
}

type WrappedSignRecoverRequest struct {
	Type    Type
	Message SignRecoverRequest
}

func (d *SignRecoverRequest) Wrap() WrappedSignRecoverRequest {
	return WrappedSignRecoverRequest{Type: d.Type(), Message: *d}
}

type SignRecoverResponse struct {
	ReturnValue ReturnValue
}

func (d *SignRecoverResponse) Type() Type {
	return TypeSignRecoverResponse
}

type WrappedSignRecoverResponse struct {
	Type    Type
	Message SignRecoverResponse
}

func (d *SignRecoverResponse) Wrap() WrappedSignRecoverResponse {
	return WrappedSignRecoverResponse{Type: d.Type(), Message: *d}
}
