package messages

type DeriveKeyRequest struct {
}

func (d *DeriveKeyRequest) Type() Type {
	return TypeDeriveKeyRequest
}

type WrappedDeriveKeyRequest struct {
	Type    Type
	Message DeriveKeyRequest
}

func (d *DeriveKeyRequest) Wrap() WrappedDeriveKeyRequest {
	return WrappedDeriveKeyRequest{Type: d.Type(), Message: *d}
}

type DeriveKeyResponse struct {
	ReturnValue ReturnValue
}

func (d *DeriveKeyResponse) Type() Type {
	return TypeDeriveKeyResponse
}

type WrappedDeriveKeyResponse struct {
	Type    Type
	Message DeriveKeyResponse
}

func (d *DeriveKeyResponse) Wrap() WrappedDeriveKeyResponse {
	return WrappedDeriveKeyResponse{Type: d.Type(), Message: *d}
}
