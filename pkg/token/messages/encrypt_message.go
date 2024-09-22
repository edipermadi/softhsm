package messages

type EncryptMessageRequest struct {
}

func (d *EncryptMessageRequest) Type() Type {
	return TypeEncryptMessageRequest
}

type WrappedEncryptMessageRequest struct {
	Type    Type
	Message EncryptMessageRequest
}

func (d *EncryptMessageRequest) Wrap() WrappedEncryptMessageRequest {
	return WrappedEncryptMessageRequest{Type: d.Type(), Message: *d}
}

type EncryptMessageResponse struct {
	ReturnValue ReturnValue
}

func (d *EncryptMessageResponse) Type() Type {
	return TypeEncryptMessageResponse
}

type WrappedEncryptMessageResponse struct {
	Type    Type
	Message EncryptMessageResponse
}

func (d *EncryptMessageResponse) Wrap() WrappedEncryptMessageResponse {
	return WrappedEncryptMessageResponse{Type: d.Type(), Message: *d}
}
