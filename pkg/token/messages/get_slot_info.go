package messages

type GetSlotInfoRequest struct {
}

func (d *GetSlotInfoRequest) Type() Type {
	return TypeGetSlotInfoRequest
}

type WrappedGetSlotInfoRequest struct {
	Type    Type
	Message GetSlotInfoRequest
}

func (d *GetSlotInfoRequest) Wrap() WrappedGetSlotInfoRequest {
	return WrappedGetSlotInfoRequest{Type: d.Type(), Message: *d}
}

type GetSlotInfoResponse struct {
	ReturnValue ReturnValue
}

func (d *GetSlotInfoResponse) Type() Type {
	return TypeGetSlotInfoResponse
}

type WrappedGetSlotInfoResponse struct {
	Type    Type
	Message GetSlotInfoResponse
}

func (d *GetSlotInfoResponse) Wrap() WrappedGetSlotInfoResponse {
	return WrappedGetSlotInfoResponse{Type: d.Type(), Message: *d}
}
