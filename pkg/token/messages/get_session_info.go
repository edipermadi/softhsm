package messages

type GetSessionInfoRequest struct {
}

func (d *GetSessionInfoRequest) Type() Type {
	return TypeGetSessionInfoRequest
}

type WrappedGetSessionInfoRequest struct {
	Type    Type
	Message GetSessionInfoRequest
}

func (d *GetSessionInfoRequest) Wrap() WrappedGetSessionInfoRequest {
	return WrappedGetSessionInfoRequest{Type: d.Type(), Message: *d}
}

type GetSessionInfoResponse struct {
	ReturnValue ReturnValue
}

func (d *GetSessionInfoResponse) Type() Type {
	return TypeGetSessionInfoResponse
}

type WrappedGetSessionInfoResponse struct {
	Type    Type
	Message GetSessionInfoResponse
}

func (d *GetSessionInfoResponse) Wrap() WrappedGetSessionInfoResponse {
	return WrappedGetSessionInfoResponse{Type: d.Type(), Message: *d}
}
