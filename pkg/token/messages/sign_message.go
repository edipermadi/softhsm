package messages

type SignMessageRequest struct {
}

func (d *SignMessageRequest) Type() Type {
	return TypeSignMessageRequest
}

type WrappedSignMessageRequest struct {
	Type    Type
	Message SignMessageRequest
}

func (d *SignMessageRequest) Wrap() WrappedSignMessageRequest {
	return WrappedSignMessageRequest{Type: d.Type(), Message: *d}
}

type SignMessageResponse struct {
	ReturnValue ReturnValue
}

func (d *SignMessageResponse) Type() Type {
	return TypeSignMessageResponse
}

type WrappedSignMessageResponse struct {
	Type    Type
	Message SignMessageResponse
}

func (d *SignMessageResponse) Wrap() WrappedSignMessageResponse {
	return WrappedSignMessageResponse{Type: d.Type(), Message: *d}
}
