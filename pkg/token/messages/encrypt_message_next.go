package messages

type EncryptMessageNextRequest struct {
}

func (d *EncryptMessageNextRequest) Type() Type {
	return TypeEncryptMessageNextRequest
}

type WrappedEncryptMessageNextRequest struct {
	Type    Type
	Message EncryptMessageNextRequest
}

func (d *EncryptMessageNextRequest) Wrap() WrappedEncryptMessageNextRequest {
	return WrappedEncryptMessageNextRequest{Type: d.Type(), Message: *d}
}

type EncryptMessageNextResponse struct {
	ReturnValue ReturnValue
}

func (d *EncryptMessageNextResponse) Type() Type {
	return TypeEncryptMessageNextResponse
}

type WrappedEncryptMessageNextResponse struct {
	Type    Type
	Message EncryptMessageNextResponse
}

func (d *EncryptMessageNextResponse) Wrap() WrappedEncryptMessageNextResponse {
	return WrappedEncryptMessageNextResponse{Type: d.Type(), Message: *d}
}
