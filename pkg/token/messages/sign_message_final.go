package messages

type SignMessageFinalRequest struct {
}

func (d *SignMessageFinalRequest) Type() Type {
	return TypeSignMessageFinalRequest
}

type WrappedSignMessageFinalRequest struct {
	Type    Type
	Message SignMessageFinalRequest
}

func (d *SignMessageFinalRequest) Wrap() WrappedSignMessageFinalRequest {
	return WrappedSignMessageFinalRequest{Type: d.Type(), Message: *d}
}

type SignMessageFinalResponse struct {
	ReturnValue ReturnValue
}

func (d *SignMessageFinalResponse) Type() Type {
	return TypeSignMessageFinalResponse
}

type WrappedSignMessageFinalResponse struct {
	Type    Type
	Message SignMessageFinalResponse
}

func (d *SignMessageFinalResponse) Wrap() WrappedSignMessageFinalResponse {
	return WrappedSignMessageFinalResponse{Type: d.Type(), Message: *d}
}
