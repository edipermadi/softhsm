package messages

type EncryptRequest struct {
}

func (d *EncryptRequest) Type() Type {
	return TypeEncryptRequest
}

type WrappedEncryptRequest struct {
	Type    Type
	Message EncryptRequest
}

func (d *EncryptRequest) Wrap() WrappedEncryptRequest {
	return WrappedEncryptRequest{Type: d.Type(), Message: *d}
}

type EncryptResponse struct {
	ReturnValue ReturnValue
}

func (d *EncryptResponse) Type() Type {
	return TypeEncryptResponse
}

type WrappedEncryptResponse struct {
	Type    Type
	Message EncryptResponse
}

func (d *EncryptResponse) Wrap() WrappedEncryptResponse {
	return WrappedEncryptResponse{Type: d.Type(), Message: *d}
}
