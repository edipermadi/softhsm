package messages

type EncryptInitRequest struct {
}

func (d *EncryptInitRequest) Type() Type {
	return TypeEncryptInitRequest
}

type WrappedEncryptInitRequest struct {
	Type    Type
	Message EncryptInitRequest
}

func (d *EncryptInitRequest) Wrap() WrappedEncryptInitRequest {
	return WrappedEncryptInitRequest{Type: d.Type(), Message: *d}
}

type EncryptInitResponse struct {
	ReturnValue ReturnValue
}

func (d *EncryptInitResponse) Type() Type {
	return TypeEncryptInitResponse
}

type WrappedEncryptInitResponse struct {
	Type    Type
	Message EncryptInitResponse
}

func (d *EncryptInitResponse) Wrap() WrappedEncryptInitResponse {
	return WrappedEncryptInitResponse{Type: d.Type(), Message: *d}
}
