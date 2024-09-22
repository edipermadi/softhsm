package messages

type SignMessageInitRequest struct {
}

func (d *SignMessageInitRequest) Type() Type {
	return TypeSignMessageInitRequest
}

type WrappedSignMessageInitRequest struct {
	Type    Type
	Message SignMessageInitRequest
}

func (d *SignMessageInitRequest) Wrap() WrappedSignMessageInitRequest {
	return WrappedSignMessageInitRequest{Type: d.Type(), Message: *d}
}

type SignMessageInitResponse struct {
	ReturnValue ReturnValue
}

func (d *SignMessageInitResponse) Type() Type {
	return TypeSignMessageInitResponse
}

type WrappedSignMessageInitResponse struct {
	Type    Type
	Message SignMessageInitResponse
}

func (d *SignMessageInitResponse) Wrap() WrappedSignMessageInitResponse {
	return WrappedSignMessageInitResponse{Type: d.Type(), Message: *d}
}
