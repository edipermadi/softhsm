package messages

type SignUpdateRequest struct {
}

func (d *SignUpdateRequest) Type() Type {
	return TypeSignUpdateRequest
}

type WrappedSignUpdateRequest struct {
	Type    Type
	Message SignUpdateRequest
}

func (d *SignUpdateRequest) Wrap() WrappedSignUpdateRequest {
	return WrappedSignUpdateRequest{Type: d.Type(), Message: *d}
}

type SignUpdateResponse struct {
	ReturnValue ReturnValue
}

func (d *SignUpdateResponse) Type() Type {
	return TypeSignUpdateResponse
}

type WrappedSignUpdateResponse struct {
	Type    Type
	Message SignUpdateResponse
}

func (d *SignUpdateResponse) Wrap() WrappedSignUpdateResponse {
	return WrappedSignUpdateResponse{Type: d.Type(), Message: *d}
}
