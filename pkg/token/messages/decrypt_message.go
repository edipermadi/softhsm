package messages

type DecryptMessageRequest struct {
}

func (d *DecryptMessageRequest) Type() Type {
	return TypeDecryptMessageRequest
}

type WrappedDecryptMessageRequest struct {
	Type    Type
	Message DecryptMessageRequest
}

func (d *DecryptMessageRequest) Wrap() WrappedDecryptMessageRequest {
	return WrappedDecryptMessageRequest{Type: d.Type(), Message: *d}
}

type DecryptMessageResponse struct {
	ReturnValue ReturnValue
}

func (d *DecryptMessageResponse) Type() Type {
	return TypeDecryptMessageResponse
}

type WrappedDecryptMessageResponse struct {
	Type    Type
	Message DecryptMessageResponse
}

func (d *DecryptMessageResponse) Wrap() WrappedDecryptMessageResponse {
	return WrappedDecryptMessageResponse{Type: d.Type(), Message: *d}
}
