package messages

type GetFunctionStatusRequest struct {
}

func (d *GetFunctionStatusRequest) Type() Type {
	return TypeGetFunctionStatusRequest
}

type WrappedGetFunctionStatusRequest struct {
	Type    Type
	Message GetFunctionStatusRequest
}

func (d *GetFunctionStatusRequest) Wrap() WrappedGetFunctionStatusRequest {
	return WrappedGetFunctionStatusRequest{Type: d.Type(), Message: *d}
}

type GetFunctionStatusResponse struct {
	ReturnValue ReturnValue
}

func (d *GetFunctionStatusResponse) Type() Type {
	return TypeGetFunctionStatusResponse
}

type WrappedGetFunctionStatusResponse struct {
	Type    Type
	Message GetFunctionStatusResponse
}

func (d *GetFunctionStatusResponse) Wrap() WrappedGetFunctionStatusResponse {
	return WrappedGetFunctionStatusResponse{Type: d.Type(), Message: *d}
}
