package messages

type DecryptInitRequest struct {
}

func (d *DecryptInitRequest) Type() Type {
	return TypeDecryptInitRequest
}

type WrappedDecryptInitRequest struct {
	Type    Type
	Message DecryptInitRequest
}

func (d *DecryptInitRequest) Wrap() WrappedDecryptInitRequest {
	return WrappedDecryptInitRequest{Type: d.Type(), Message: *d}
}

type DecryptInitResponse struct {
	ReturnValue ReturnValue
}

func (d *DecryptInitResponse) Type() Type {
	return TypeDecryptInitResponse
}

type WrappedDecryptInitResponse struct {
	Type    Type
	Message DecryptInitResponse
}

func (d *DecryptInitResponse) Wrap() WrappedDecryptInitResponse {
	return WrappedDecryptInitResponse{Type: d.Type(), Message: *d}
}
