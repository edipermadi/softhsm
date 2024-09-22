package messages

type DecryptMessageInitRequest struct {
}

func (d *DecryptMessageInitRequest) Type() Type {
	return TypeDecryptMessageInitRequest
}

type WrappedDecryptMessageInitRequest struct {
	Type    Type
	Message DecryptMessageInitRequest
}

func (d *DecryptMessageInitRequest) Wrap() WrappedDecryptMessageInitRequest {
	return WrappedDecryptMessageInitRequest{Type: d.Type(), Message: *d}
}

type DecryptMessageInitResponse struct {
	ReturnValue ReturnValue
}

func (d *DecryptMessageInitResponse) Type() Type {
	return TypeDecryptMessageInitResponse
}

type WrappedDecryptMessageInitResponse struct {
	Type    Type
	Message DecryptMessageInitResponse
}

func (d *DecryptMessageInitResponse) Wrap() WrappedDecryptMessageInitResponse {
	return WrappedDecryptMessageInitResponse{Type: d.Type(), Message: *d}
}
