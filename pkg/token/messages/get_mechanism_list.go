package messages

type GetMechanismListRequest struct {
}

func (d *GetMechanismListRequest) Type() Type {
	return TypeGetMechanismListRequest
}

type WrappedGetMechanismListRequest struct {
	Type    Type
	Message GetMechanismListRequest
}

func (d *GetMechanismListRequest) Wrap() WrappedGetMechanismListRequest {
	return WrappedGetMechanismListRequest{Type: d.Type(), Message: *d}
}

type GetMechanismListResponse struct {
	ReturnValue ReturnValue
}

func (d *GetMechanismListResponse) Type() Type {
	return TypeGetMechanismListResponse
}

type WrappedGetMechanismListResponse struct {
	Type    Type
	Message GetMechanismListResponse
}

func (d *GetMechanismListResponse) Wrap() WrappedGetMechanismListResponse {
	return WrappedGetMechanismListResponse{Type: d.Type(), Message: *d}
}
