package messages

type EncryptUpdateRequest struct {
}

func (d *EncryptUpdateRequest) Type() Type {
	return TypeEncryptUpdateRequest
}

type WrappedEncryptUpdateRequest struct {
	Type    Type
	Message EncryptUpdateRequest
}

func (d *EncryptUpdateRequest) Wrap() WrappedEncryptUpdateRequest {
	return WrappedEncryptUpdateRequest{Type: d.Type(), Message: *d}
}

type EncryptUpdateResponse struct {
	ReturnValue ReturnValue
}

func (d *EncryptUpdateResponse) Type() Type {
	return TypeEncryptUpdateResponse
}

type WrappedEncryptUpdateResponse struct {
	Type    Type
	Message EncryptUpdateResponse
}

func (d *EncryptUpdateResponse) Wrap() WrappedEncryptUpdateResponse {
	return WrappedEncryptUpdateResponse{Type: d.Type(), Message: *d}
}
