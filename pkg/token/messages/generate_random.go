package messages

type GenerateRandomRequest struct {
	SessionID int
}

func (d *GenerateRandomRequest) Type() Type {
	return TypeGenerateRandomRequest
}

type WrappedGenerateRandomRequest struct {
	Type    Type
	Message GenerateRandomRequest
}

func (d *GenerateRandomRequest) Wrap() WrappedGenerateRandomRequest {
	return WrappedGenerateRandomRequest{Type: d.Type(), Message: *d}
}

type GenerateRandomResponse struct {
	ReturnValue ReturnValue
	Random      []byte
}

func (d *GenerateRandomResponse) Type() Type {
	return TypeGenerateRandomResponse
}

type WrappedGenerateRandomResponse struct {
	Type    Type
	Message GenerateRandomResponse
}

func (d *GenerateRandomResponse) Wrap() WrappedGenerateRandomResponse {
	return WrappedGenerateRandomResponse{Type: d.Type(), Message: *d}
}
