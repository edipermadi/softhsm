package messages

type GetFunctionListRequest struct {
}

func (d *GetFunctionListRequest) Type() Type {
	return TypeGetFunctionListRequest
}

type WrappedGetFunctionListRequest struct {
	Type    Type
	Message GetFunctionListRequest
}

func (d *GetFunctionListRequest) Wrap() WrappedGetFunctionListRequest {
	return WrappedGetFunctionListRequest{Type: d.Type(), Message: *d}
}

type GetFunctionListResponse struct {
	ReturnValue ReturnValue
}

func (d *GetFunctionListResponse) Type() Type {
	return TypeGetFunctionListResponse
}

type WrappedGetFunctionListResponse struct {
	Type    Type
	Message GetFunctionListResponse
}

func (d *GetFunctionListResponse) Wrap() WrappedGetFunctionListResponse {
	return WrappedGetFunctionListResponse{Type: d.Type(), Message: *d}
}
