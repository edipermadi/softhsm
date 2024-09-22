package messages

type DecryptUpdateRequest struct {
}

func (d *DecryptUpdateRequest) Type() Type {
	return TypeDecryptUpdateRequest
}

type WrappedDecryptUpdateRequest struct {
	Type    Type
	Message DecryptUpdateRequest
}

func (d *DecryptUpdateRequest) Wrap() WrappedDecryptUpdateRequest {
	return WrappedDecryptUpdateRequest{Type: d.Type(), Message: *d}
}

type DecryptUpdateResponse struct {
	ReturnValue ReturnValue
}

func (d *DecryptUpdateResponse) Type() Type {
	return TypeDecryptUpdateResponse
}

type WrappedDecryptUpdateResponse struct {
	Type    Type
	Message DecryptUpdateResponse
}

func (d *DecryptUpdateResponse) Wrap() WrappedDecryptUpdateResponse {
	return WrappedDecryptUpdateResponse{Type: d.Type(), Message: *d}
}
