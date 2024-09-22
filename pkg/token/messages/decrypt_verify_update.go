package messages

type DecryptVerifyUpdateRequest struct {
}

func (d *DecryptVerifyUpdateRequest) Type() Type {
	return TypeDecryptVerifyUpdateRequest
}

type WrappedDecryptVerifyUpdateRequest struct {
	Type    Type
	Message DecryptVerifyUpdateRequest
}

func (d *DecryptVerifyUpdateRequest) Wrap() WrappedDecryptVerifyUpdateRequest {
	return WrappedDecryptVerifyUpdateRequest{Type: d.Type(), Message: *d}
}

type DecryptVerifyUpdateResponse struct {
	ReturnValue ReturnValue
}

func (d *DecryptVerifyUpdateResponse) Type() Type {
	return TypeDecryptVerifyUpdateResponse
}

type WrappedDecryptVerifyUpdateResponse struct {
	Type    Type
	Message DecryptVerifyUpdateResponse
}

func (d *DecryptVerifyUpdateResponse) Wrap() WrappedDecryptVerifyUpdateResponse {
	return WrappedDecryptVerifyUpdateResponse{Type: d.Type(), Message: *d}
}
