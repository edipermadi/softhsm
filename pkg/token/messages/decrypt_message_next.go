package messages

type DecryptMessageNextRequest struct {
}

func (d *DecryptMessageNextRequest) Type() Type {
	return TypeDecryptMessageNextRequest
}

type WrappedDecryptMessageNextRequest struct {
	Type    Type
	Message DecryptMessageNextRequest
}

func (d *DecryptMessageNextRequest) Wrap() WrappedDecryptMessageNextRequest {
	return WrappedDecryptMessageNextRequest{Type: d.Type(), Message: *d}
}

type DecryptMessageNextResponse struct {
	ReturnValue ReturnValue
}

func (d *DecryptMessageNextResponse) Type() Type {
	return TypeDecryptMessageNextResponse
}

type WrappedDecryptMessageNextResponse struct {
	Type    Type
	Message DecryptMessageNextResponse
}

func (d *DecryptMessageNextResponse) Wrap() WrappedDecryptMessageNextResponse {
	return WrappedDecryptMessageNextResponse{Type: d.Type(), Message: *d}
}
