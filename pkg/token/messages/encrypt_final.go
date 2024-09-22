package messages

type EncryptFinalRequest struct {
}

func (d *EncryptFinalRequest) Type() Type {
	return TypeEncryptFinalRequest
}

type WrappedEncryptFinalRequest struct {
	Type    Type
	Message EncryptFinalRequest
}

func (d *EncryptFinalRequest) Wrap() WrappedEncryptFinalRequest {
	return WrappedEncryptFinalRequest{Type: d.Type(), Message: *d}
}

type EncryptFinalResponse struct {
	ReturnValue ReturnValue
}

func (d *EncryptFinalResponse) Type() Type {
	return TypeEncryptFinalResponse
}

type WrappedEncryptFinalResponse struct {
	Type    Type
	Message EncryptFinalResponse
}

func (d *EncryptFinalResponse) Wrap() WrappedEncryptFinalResponse {
	return WrappedEncryptFinalResponse{Type: d.Type(), Message: *d}
}
