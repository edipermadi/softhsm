package messages

type GetInfoRequest struct {
}

func (d *GetInfoRequest) Type() Type {
	return TypeGetInfoRequest
}

type WrappedGetInfoRequest struct {
	Type    Type
	Message GetInfoRequest
}

func (d *GetInfoRequest) Wrap() WrappedGetInfoRequest {
	return WrappedGetInfoRequest{Type: d.Type(), Message: *d}
}

type GetInfoResponse struct {
	ReturnValue ReturnValue
}

func (d *GetInfoResponse) Type() Type {
	return TypeGetInfoResponse
}

type WrappedGetInfoResponse struct {
	Type    Type
	Message GetInfoResponse
}

func (d *GetInfoResponse) Wrap() WrappedGetInfoResponse {
	return WrappedGetInfoResponse{Type: d.Type(), Message: *d}
}
