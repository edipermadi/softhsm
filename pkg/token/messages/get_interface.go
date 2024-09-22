package messages

type GetInterfaceRequest struct {
}

func (d *GetInterfaceRequest) Type() Type {
	return TypeGetInterfaceRequest
}

type WrappedGetInterfaceRequest struct {
	Type    Type
	Message GetInterfaceRequest
}

func (d *GetInterfaceRequest) Wrap() WrappedGetInterfaceRequest {
	return WrappedGetInterfaceRequest{Type: d.Type(), Message: *d}
}

type GetInterfaceResponse struct {
	ReturnValue ReturnValue
}

func (d *GetInterfaceResponse) Type() Type {
	return TypeGetInterfaceResponse
}

type WrappedGetInterfaceResponse struct {
	Type    Type
	Message GetInterfaceResponse
}

func (d *GetInterfaceResponse) Wrap() WrappedGetInterfaceResponse {
	return WrappedGetInterfaceResponse{Type: d.Type(), Message: *d}
}
