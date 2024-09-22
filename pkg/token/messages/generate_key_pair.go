package messages

type GenerateKeyPairRequest struct {
}

func (d *GenerateKeyPairRequest) Type() Type {
	return TypeGenerateKeyPairRequest
}

type WrappedGenerateKeyPairRequest struct {
	Type    Type
	Message GenerateKeyPairRequest
}

func (d *GenerateKeyPairRequest) Wrap() WrappedGenerateKeyPairRequest {
	return WrappedGenerateKeyPairRequest{Type: d.Type(), Message: *d}
}

type GenerateKeyPairResponse struct {
	ReturnValue ReturnValue
}

func (d *GenerateKeyPairResponse) Type() Type {
	return TypeGenerateKeyPairResponse
}

type WrappedGenerateKeyPairResponse struct {
	Type    Type
	Message GenerateKeyPairResponse
}

func (d *GenerateKeyPairResponse) Wrap() WrappedGenerateKeyPairResponse {
	return WrappedGenerateKeyPairResponse{Type: d.Type(), Message: *d}
}
