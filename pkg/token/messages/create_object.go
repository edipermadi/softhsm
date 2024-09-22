package messages

type CreateObjectRequest struct {
}

func (d *CreateObjectRequest) Type() Type {
	return TypeCreateObjectRequest
}

type WrappedCreateObjectRequest struct {
	Type    Type
	Message CreateObjectRequest
}

func (d *CreateObjectRequest) Wrap() WrappedCreateObjectRequest {
	return WrappedCreateObjectRequest{Type: d.Type(), Message: *d}
}

type CreateObjectResponse struct {
	ReturnValue ReturnValue
}

func (d *CreateObjectResponse) Type() Type {
	return TypeCreateObjectResponse
}

type WrappedCreateObjectResponse struct {
	Type    Type
	Message CreateObjectResponse
}

func (d *CreateObjectResponse) Wrap() WrappedCreateObjectResponse {
	return WrappedCreateObjectResponse{Type: d.Type(), Message: *d}
}
