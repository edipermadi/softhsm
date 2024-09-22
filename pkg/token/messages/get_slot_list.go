package messages

type GetSlotListRequest struct {
}

func (d *GetSlotListRequest) Type() Type {
	return TypeGetSlotListRequest
}

type WrappedGetSlotListRequest struct {
	Type    Type
	Message GetSlotListRequest
}

func (d *GetSlotListRequest) Wrap() WrappedGetSlotListRequest {
	return WrappedGetSlotListRequest{Type: d.Type(), Message: *d}
}

type GetSlotListResponse struct {
	ReturnValue ReturnValue
}

func (d *GetSlotListResponse) Type() Type {
	return TypeGetSlotListResponse
}

type WrappedGetSlotListResponse struct {
	Type    Type
	Message GetSlotListResponse
}

func (d *GetSlotListResponse) Wrap() WrappedGetSlotListResponse {
	return WrappedGetSlotListResponse{Type: d.Type(), Message: *d}
}
