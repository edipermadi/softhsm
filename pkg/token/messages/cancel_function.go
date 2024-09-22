package messages

type CancelFunctionRequest struct {
}

func (d *CancelFunctionRequest) Type() Type {
	return TypeCancelFunctionRequest
}

type WrappedCancelFunctionRequest struct {
	Type    Type
	Message CancelFunctionRequest
}

func (d *CancelFunctionRequest) Wrap() WrappedCancelFunctionRequest {
	return WrappedCancelFunctionRequest{Type: d.Type(), Message: *d}
}

type CancelFunctionResponse struct {
	ReturnValue ReturnValue
}

func (d *CancelFunctionResponse) Type() Type {
	return TypeCancelFunctionResponse
}

type WrappedCancelFunctionResponse struct {
	Type    Type
	Message CancelFunctionResponse
}

func (d *CancelFunctionResponse) Wrap() WrappedCancelFunctionResponse {
	return WrappedCancelFunctionResponse{Type: d.Type(), Message: *d}
}
