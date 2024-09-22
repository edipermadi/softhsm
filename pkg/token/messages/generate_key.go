package messages

type GenerateKeyRequest struct {
}

func (d *GenerateKeyRequest) Type() Type {
	return TypeGenerateKeyRequest
}

type WrappedGenerateKeyRequest struct {
	Type    Type
	Message GenerateKeyRequest
}

func (d *GenerateKeyRequest) Wrap() WrappedGenerateKeyRequest {
	return WrappedGenerateKeyRequest{Type: d.Type(), Message: *d}
}

type GenerateKeyResponse struct {
	ReturnValue ReturnValue
}

func (d *GenerateKeyResponse) Type() Type {
	return TypeGenerateKeyResponse
}

type WrappedGenerateKeyResponse struct {
	Type    Type
	Message GenerateKeyResponse
}

func (d *GenerateKeyResponse) Wrap() WrappedGenerateKeyResponse {
	return WrappedGenerateKeyResponse{Type: d.Type(), Message: *d}
}
