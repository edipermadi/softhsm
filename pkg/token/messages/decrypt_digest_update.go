package messages

type DecryptDigestUpdateRequest struct {
}

func (d *DecryptDigestUpdateRequest) Type() Type {
	return TypeDecryptDigestUpdateRequest
}

type WrappedDecryptDigestUpdateRequest struct {
	Type    Type
	Message DecryptDigestUpdateRequest
}

func (d *DecryptDigestUpdateRequest) Wrap() WrappedDecryptDigestUpdateRequest {
	return WrappedDecryptDigestUpdateRequest{Type: d.Type(), Message: *d}
}

type DecryptDigestUpdateResponse struct {
	ReturnValue ReturnValue
}

func (d *DecryptDigestUpdateResponse) Type() Type {
	return TypeDecryptDigestUpdateResponse
}

type WrappedDecryptDigestUpdateResponse struct {
	Type    Type
	Message DecryptDigestUpdateResponse
}

func (d *DecryptDigestUpdateResponse) Wrap() WrappedDecryptDigestUpdateResponse {
	return WrappedDecryptDigestUpdateResponse{Type: d.Type(), Message: *d}
}
