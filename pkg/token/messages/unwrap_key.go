package messages

type UnwrapKeyRequest struct {
}

func (d *UnwrapKeyRequest) Type() Type {
	return TypeUnwrapKeyRequest
}

type WrappedUnwrapKeyRequest struct {
	Type    Type
	Message UnwrapKeyRequest
}

func (d *UnwrapKeyRequest) Wrap() WrappedUnwrapKeyRequest {
	return WrappedUnwrapKeyRequest{Type: d.Type(), Message: *d}
}

type UnwrapKeyResponse struct {
	ReturnValue ReturnValue
}

func (d *UnwrapKeyResponse) Type() Type {
	return TypeUnwrapKeyResponse
}

type WrappedUnwrapKeyResponse struct {
	Type    Type
	Message UnwrapKeyResponse
}

func (d *UnwrapKeyResponse) Wrap() WrappedUnwrapKeyResponse {
	return WrappedUnwrapKeyResponse{Type: d.Type(), Message: *d}
}
