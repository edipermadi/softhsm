package messages

type GetTokenInfoRequest struct {
}

func (d *GetTokenInfoRequest) Type() Type {
	return TypeGetTokenInfoRequest
}

type WrappedGetTokenInfoRequest struct {
	Type    Type
	Message GetTokenInfoRequest
}

func (d *GetTokenInfoRequest) Wrap() WrappedGetTokenInfoRequest {
	return WrappedGetTokenInfoRequest{Type: d.Type(), Message: *d}
}

type GetTokenInfoResponse struct {
	ReturnValue ReturnValue
}

func (d *GetTokenInfoResponse) Type() Type {
	return TypeGetTokenInfoResponse
}

type WrappedGetTokenInfoResponse struct {
	Type    Type
	Message GetTokenInfoResponse
}

func (d *GetTokenInfoResponse) Wrap() WrappedGetTokenInfoResponse {
	return WrappedGetTokenInfoResponse{Type: d.Type(), Message: *d}
}
