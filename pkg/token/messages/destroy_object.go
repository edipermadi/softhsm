package messages

type DestroyObjectRequest struct {
}

func (d *DestroyObjectRequest) Type() Type {
	return TypeDestroyObjectRequest
}

type WrappedDestroyObjectRequest struct {
	Type    Type
	Message DestroyObjectRequest
}

func (d *DestroyObjectRequest) Wrap() WrappedDestroyObjectRequest {
	return WrappedDestroyObjectRequest{Type: d.Type(), Message: *d}
}

type DestroyObjectResponse struct {
	ReturnValue ReturnValue
}

func (d *DestroyObjectResponse) Type() Type {
	return TypeDestroyObjectResponse
}

type WrappedDestroyObjectResponse struct {
	Type    Type
	Message DestroyObjectResponse
}

func (d *DestroyObjectResponse) Wrap() WrappedDestroyObjectResponse {
	return WrappedDestroyObjectResponse{Type: d.Type(), Message: *d}
}
