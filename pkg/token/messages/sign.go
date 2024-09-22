package messages

type SignRequest struct {
}

func (d *SignRequest) Type() Type {
	return TypeSignRequest
}

type WrappedSignRequest struct {
	Type    Type
	Message SignRequest
}

func (d *SignRequest) Wrap() WrappedSignRequest {
	return WrappedSignRequest{Type: d.Type(), Message: *d}
}

type SignResponse struct {
	ReturnValue ReturnValue
}

func (d *SignResponse) Type() Type {
	return TypeSignResponse
}

type WrappedSignResponse struct {
	Type    Type
	Message SignResponse
}

func (d *SignResponse) Wrap() WrappedSignResponse {
	return WrappedSignResponse{Type: d.Type(), Message: *d}
}
