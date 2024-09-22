package messages

type VerifyUpdateRequest struct {
}

func (d *VerifyUpdateRequest) Type() Type {
	return TypeVerifyUpdateRequest
}

type WrappedVerifyUpdateRequest struct {
	Type    Type
	Message VerifyUpdateRequest
}

func (d *VerifyUpdateRequest) Wrap() WrappedVerifyUpdateRequest {
	return WrappedVerifyUpdateRequest{Type: d.Type(), Message: *d}
}

type VerifyUpdateResponse struct {
	ReturnValue ReturnValue
}

func (d *VerifyUpdateResponse) Type() Type {
	return TypeVerifyUpdateResponse
}

type WrappedVerifyUpdateResponse struct {
	Type    Type
	Message VerifyUpdateResponse
}

func (d *VerifyUpdateResponse) Wrap() WrappedVerifyUpdateResponse {
	return WrappedVerifyUpdateResponse{Type: d.Type(), Message: *d}
}
