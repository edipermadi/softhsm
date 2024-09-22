package messages

type CopyObjectRequest struct {
}

func (d *CopyObjectRequest) Type() Type {
	return TypeCopyObjectRequest
}

type WrappedCopyObjectRequest struct {
	Type    Type
	Message CopyObjectRequest
}

func (d *CopyObjectRequest) Wrap() WrappedCopyObjectRequest {
	return WrappedCopyObjectRequest{Type: d.Type(), Message: *d}
}

type CopyObjectResponse struct {
	ReturnValue ReturnValue
}

func (d *CopyObjectResponse) Type() Type {
	return TypeCopyObjectResponse
}

type WrappedCopyObjectResponse struct {
	Type    Type
	Message CopyObjectResponse
}

func (d *CopyObjectResponse) Wrap() WrappedCopyObjectResponse {
	return WrappedCopyObjectResponse{Type: d.Type(), Message: *d}
}
