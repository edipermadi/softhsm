package messages

type DecryptMessageFinalRequest struct {
}

func (d *DecryptMessageFinalRequest) Type() Type {
	return TypeDecryptMessageFinalRequest
}

type WrappedDecryptMessageFinalRequest struct {
	Type    Type
	Message DecryptMessageFinalRequest
}

func (d *DecryptMessageFinalRequest) Wrap() WrappedDecryptMessageFinalRequest {
	return WrappedDecryptMessageFinalRequest{Type: d.Type(), Message: *d}
}

type DecryptMessageFinalResponse struct {
	ReturnValue ReturnValue
}

func (d *DecryptMessageFinalResponse) Type() Type {
	return TypeDecryptMessageFinalResponse
}

type WrappedDecryptMessageFinalResponse struct {
	Type    Type
	Message DecryptMessageFinalResponse
}

func (d *DecryptMessageFinalResponse) Wrap() WrappedDecryptMessageFinalResponse {
	return WrappedDecryptMessageFinalResponse{Type: d.Type(), Message: *d}
}
