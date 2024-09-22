package messages

type EncryptMessageInitRequest struct {
}

func (d *EncryptMessageInitRequest) Type() Type {
	return TypeEncryptMessageInitRequest
}

type WrappedEncryptMessageInitRequest struct {
	Type    Type
	Message EncryptMessageInitRequest
}

func (d *EncryptMessageInitRequest) Wrap() WrappedEncryptMessageInitRequest {
	return WrappedEncryptMessageInitRequest{Type: d.Type(), Message: *d}
}

type EncryptMessageInitResponse struct {
	ReturnValue ReturnValue
}

func (d *EncryptMessageInitResponse) Type() Type {
	return TypeEncryptMessageInitResponse
}

type WrappedEncryptMessageInitResponse struct {
	Type    Type
	Message EncryptMessageInitResponse
}

func (d *EncryptMessageInitResponse) Wrap() WrappedEncryptMessageInitResponse {
	return WrappedEncryptMessageInitResponse{Type: d.Type(), Message: *d}
}
