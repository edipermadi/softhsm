package messages

type DecryptRequest struct {
}

func (d *DecryptRequest) Type() Type {
	return TypeDecryptRequest
}

type WrappedDecryptRequest struct {
	Type    Type
	Message DecryptRequest
}

func (d *DecryptRequest) Wrap() WrappedDecryptRequest {
	return WrappedDecryptRequest{Type: d.Type(), Message: *d}
}

type DecryptResponse struct {
	ReturnValue ReturnValue
}

func (d *DecryptResponse) Type() Type {
	return TypeDecryptResponse
}

type WrappedDecryptResponse struct {
	Type    Type
	Message DecryptResponse
}

func (d *DecryptResponse) Wrap() WrappedDecryptResponse {
	return WrappedDecryptResponse{Type: d.Type(), Message: *d}
}
