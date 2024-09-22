package messages

type EncryptMessageFinalRequest struct {
}

func (d *EncryptMessageFinalRequest) Type() Type {
	return TypeEncryptMessageFinalRequest
}

type WrappedEncryptMessageFinalRequest struct {
	Type    Type
	Message EncryptMessageFinalRequest
}

func (d *EncryptMessageFinalRequest) Wrap() WrappedEncryptMessageFinalRequest {
	return WrappedEncryptMessageFinalRequest{Type: d.Type(), Message: *d}
}

type EncryptMessageFinalResponse struct {
	ReturnValue ReturnValue
}

func (d *EncryptMessageFinalResponse) Type() Type {
	return TypeEncryptMessageFinalResponse
}

type WrappedEncryptMessageFinalResponse struct {
	Type    Type
	Message EncryptMessageFinalResponse
}

func (d *EncryptMessageFinalResponse) Wrap() WrappedEncryptMessageFinalResponse {
	return WrappedEncryptMessageFinalResponse{Type: d.Type(), Message: *d}
}
