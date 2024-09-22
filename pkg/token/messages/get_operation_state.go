package messages

type GetOperationStateRequest struct {
}

func (d *GetOperationStateRequest) Type() Type {
	return TypeGetOperationStateRequest
}

type WrappedGetOperationStateRequest struct {
	Type    Type
	Message GetOperationStateRequest
}

func (d *GetOperationStateRequest) Wrap() WrappedGetOperationStateRequest {
	return WrappedGetOperationStateRequest{Type: d.Type(), Message: *d}
}

type GetOperationStateResponse struct {
	ReturnValue ReturnValue
}

func (d *GetOperationStateResponse) Type() Type {
	return TypeGetOperationStateResponse
}

type WrappedGetOperationStateResponse struct {
	Type    Type
	Message GetOperationStateResponse
}

func (d *GetOperationStateResponse) Wrap() WrappedGetOperationStateResponse {
	return WrappedGetOperationStateResponse{Type: d.Type(), Message: *d}
}
