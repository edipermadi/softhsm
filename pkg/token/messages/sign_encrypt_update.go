package messages

type SignEncryptUpdateRequest struct {
}

func (d *SignEncryptUpdateRequest) Type() Type {
	return TypeSignEncryptUpdateRequest
}

type WrappedSignEncryptUpdateRequest struct {
	Type    Type
	Message SignEncryptUpdateRequest
}

func (d *SignEncryptUpdateRequest) Wrap() WrappedSignEncryptUpdateRequest {
	return WrappedSignEncryptUpdateRequest{Type: d.Type(), Message: *d}
}

type SignEncryptUpdateResponse struct {
	ReturnValue ReturnValue
}

func (d *SignEncryptUpdateResponse) Type() Type {
	return TypeSignEncryptUpdateResponse
}

type WrappedSignEncryptUpdateResponse struct {
	Type    Type
	Message SignEncryptUpdateResponse
}

func (d *SignEncryptUpdateResponse) Wrap() WrappedSignEncryptUpdateResponse {
	return WrappedSignEncryptUpdateResponse{Type: d.Type(), Message: *d}
}
