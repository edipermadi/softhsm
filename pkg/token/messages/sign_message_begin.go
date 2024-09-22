package messages

type SignMessageBeginRequest struct {
}

func (d *SignMessageBeginRequest) Type() Type {
	return TypeSignMessageBeginRequest
}

type WrappedSignMessageBeginRequest struct {
	Type    Type
	Message SignMessageBeginRequest
}

func (d *SignMessageBeginRequest) Wrap() WrappedSignMessageBeginRequest {
	return WrappedSignMessageBeginRequest{Type: d.Type(), Message: *d}
}

type SignMessageBeginResponse struct {
	ReturnValue ReturnValue
}

func (d *SignMessageBeginResponse) Type() Type {
	return TypeSignMessageBeginResponse
}

type WrappedSignMessageBeginResponse struct {
	Type    Type
	Message SignMessageBeginResponse
}

func (d *SignMessageBeginResponse) Wrap() WrappedSignMessageBeginResponse {
	return WrappedSignMessageBeginResponse{Type: d.Type(), Message: *d}
}
