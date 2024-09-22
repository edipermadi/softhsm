package messages

type SignFinalRequest struct {
}

func (d *SignFinalRequest) Type() Type {
	return TypeSignFinalRequest
}

type WrappedSignFinalRequest struct {
	Type    Type
	Message SignFinalRequest
}

func (d *SignFinalRequest) Wrap() WrappedSignFinalRequest {
	return WrappedSignFinalRequest{Type: d.Type(), Message: *d}
}

type SignFinalResponse struct {
	ReturnValue ReturnValue
}

func (d *SignFinalResponse) Type() Type {
	return TypeSignFinalResponse
}

type WrappedSignFinalResponse struct {
	Type    Type
	Message SignFinalResponse
}

func (d *SignFinalResponse) Wrap() WrappedSignFinalResponse {
	return WrappedSignFinalResponse{Type: d.Type(), Message: *d}
}
