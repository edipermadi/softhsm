package messages

type GetInterfaceListRequest struct {
}

func (d *GetInterfaceListRequest) Type() Type {
	return TypeGetInterfaceListRequest
}

type WrappedGetInterfaceListRequest struct {
	Type    Type
	Message GetInterfaceListRequest
}

func (d *GetInterfaceListRequest) Wrap() WrappedGetInterfaceListRequest {
	return WrappedGetInterfaceListRequest{Type: d.Type(), Message: *d}
}

type GetInterfaceListResponse struct {
	ReturnValue ReturnValue
}

func (d *GetInterfaceListResponse) Type() Type {
	return TypeGetInterfaceListResponse
}

type WrappedGetInterfaceListResponse struct {
	Type    Type
	Message GetInterfaceListResponse
}

func (d *GetInterfaceListResponse) Wrap() WrappedGetInterfaceListResponse {
	return WrappedGetInterfaceListResponse{Type: d.Type(), Message: *d}
}
