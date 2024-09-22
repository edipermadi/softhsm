package messages

type DecryptFinalRequest struct {
}

func (d *DecryptFinalRequest) Type() Type {
	return TypeDecryptFinalRequest
}

type WrappedDecryptFinalRequest struct {
	Type    Type
	Message DecryptFinalRequest
}

func (d *DecryptFinalRequest) Wrap() WrappedDecryptFinalRequest {
	return WrappedDecryptFinalRequest{Type: d.Type(), Message: *d}
}

type DecryptFinalResponse struct {
	ReturnValue ReturnValue
}

func (d *DecryptFinalResponse) Type() Type {
	return TypeDecryptFinalResponse
}

type WrappedDecryptFinalResponse struct {
	Type    Type
	Message DecryptFinalResponse
}

func (d *DecryptFinalResponse) Wrap() WrappedDecryptFinalResponse {
	return WrappedDecryptFinalResponse{Type: d.Type(), Message: *d}
}
