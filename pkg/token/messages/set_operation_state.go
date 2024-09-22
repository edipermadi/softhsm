package messages

type SetOperationStateRequest struct {
}

func (d *SetOperationStateRequest) Type() Type {
	return TypeSetOperationStateRequest
}

type WrappedSetOperationStateRequest struct {
	Type    Type
	Message SetOperationStateRequest
}

func (d *SetOperationStateRequest) Wrap() WrappedSetOperationStateRequest {
	return WrappedSetOperationStateRequest{Type: d.Type(), Message: *d}
}

type SetOperationStateResponse struct {
	ReturnValue ReturnValue
}

func (d *SetOperationStateResponse) Type() Type {
	return TypeSetOperationStateResponse
}

type WrappedSetOperationStateResponse struct {
	Type    Type
	Message SetOperationStateResponse
}

func (d *SetOperationStateResponse) Wrap() WrappedSetOperationStateResponse {
	return WrappedSetOperationStateResponse{Type: d.Type(), Message: *d}
}
