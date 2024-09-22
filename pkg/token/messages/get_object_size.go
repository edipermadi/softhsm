package messages

type GetObjectSizeRequest struct {
}

func (d *GetObjectSizeRequest) Type() Type {
	return TypeGetObjectSizeRequest
}

type WrappedGetObjectSizeRequest struct {
	Type    Type
	Message GetObjectSizeRequest
}

func (d *GetObjectSizeRequest) Wrap() WrappedGetObjectSizeRequest {
	return WrappedGetObjectSizeRequest{Type: d.Type(), Message: *d}
}

type GetObjectSizeResponse struct {
	ReturnValue ReturnValue
}

func (d *GetObjectSizeResponse) Type() Type {
	return TypeGetObjectSizeResponse
}

type WrappedGetObjectSizeResponse struct {
	Type    Type
	Message GetObjectSizeResponse
}

func (d *GetObjectSizeResponse) Wrap() WrappedGetObjectSizeResponse {
	return WrappedGetObjectSizeResponse{Type: d.Type(), Message: *d}
}
