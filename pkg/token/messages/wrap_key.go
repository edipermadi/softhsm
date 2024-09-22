package messages

type WrapKeyRequest struct {
}

func (d *WrapKeyRequest) Type() Type {
	return TypeWrapKeyRequest
}

type WrappedWrapKeyRequest struct {
	Type    Type
	Message WrapKeyRequest
}

func (d *WrapKeyRequest) Wrap() WrappedWrapKeyRequest {
	return WrappedWrapKeyRequest{Type: d.Type(), Message: *d}
}

type WrapKeyResponse struct {
	ReturnValue ReturnValue
}

func (d *WrapKeyResponse) Type() Type {
	return TypeWrapKeyResponse
}

type WrappedWrapKeyResponse struct {
	Type    Type
	Message WrapKeyResponse
}

func (d *WrapKeyResponse) Wrap() WrappedWrapKeyResponse {
	return WrappedWrapKeyResponse{Type: d.Type(), Message: *d}
}
