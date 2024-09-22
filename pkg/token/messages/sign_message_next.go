package messages

type SignMessageNextRequest struct {
}

func (d *SignMessageNextRequest) Type() Type {
	return TypeSignMessageNextRequest
}

type WrappedSignMessageNextRequest struct {
	Type    Type
	Message SignMessageNextRequest
}

func (d *SignMessageNextRequest) Wrap() WrappedSignMessageNextRequest {
	return WrappedSignMessageNextRequest{Type: d.Type(), Message: *d}
}

type SignMessageNextResponse struct {
	ReturnValue ReturnValue
}

func (d *SignMessageNextResponse) Type() Type {
	return TypeSignMessageNextResponse
}

type WrappedSignMessageNextResponse struct {
	Type    Type
	Message SignMessageNextResponse
}

func (d *SignMessageNextResponse) Wrap() WrappedSignMessageNextResponse {
	return WrappedSignMessageNextResponse{Type: d.Type(), Message: *d}
}
