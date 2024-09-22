package messages

type SignInitRequest struct {
}

func (d *SignInitRequest) Type() Type {
	return TypeSignInitRequest
}

type WrappedSignInitRequest struct {
	Type    Type
	Message SignInitRequest
}

func (d *SignInitRequest) Wrap() WrappedSignInitRequest {
	return WrappedSignInitRequest{Type: d.Type(), Message: *d}
}

type SignInitResponse struct {
	ReturnValue ReturnValue
}

func (d *SignInitResponse) Type() Type {
	return TypeSignInitResponse
}

type WrappedSignInitResponse struct {
	Type    Type
	Message SignInitResponse
}

func (d *SignInitResponse) Wrap() WrappedSignInitResponse {
	return WrappedSignInitResponse{Type: d.Type(), Message: *d}
}
