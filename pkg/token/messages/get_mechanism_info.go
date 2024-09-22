package messages

type GetMechanismInfoRequest struct {
}

func (d *GetMechanismInfoRequest) Type() Type {
	return TypeGetMechanismInfoRequest
}

type WrappedGetMechanismInfoRequest struct {
	Type    Type
	Message GetMechanismInfoRequest
}

func (d *GetMechanismInfoRequest) Wrap() WrappedGetMechanismInfoRequest {
	return WrappedGetMechanismInfoRequest{Type: d.Type(), Message: *d}
}

type GetMechanismInfoResponse struct {
	ReturnValue ReturnValue
}

func (d *GetMechanismInfoResponse) Type() Type {
	return TypeGetMechanismInfoResponse
}

type WrappedGetMechanismInfoResponse struct {
	Type    Type
	Message GetMechanismInfoResponse
}

func (d *GetMechanismInfoResponse) Wrap() WrappedGetMechanismInfoResponse {
	return WrappedGetMechanismInfoResponse{Type: d.Type(), Message: *d}
}
