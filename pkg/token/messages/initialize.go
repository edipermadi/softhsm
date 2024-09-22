package messages

type InitializeRequest struct {
}

func (d *InitializeRequest) Type() Type {
	return TypeInitializeRequest
}

type WrappedInitializeRequest struct {
	Type    Type
	Message InitializeRequest
}

func (d *InitializeRequest) Wrap() WrappedInitializeRequest {
	return WrappedInitializeRequest{Type: d.Type(), Message: *d}
}

type InitializeResponse struct {
	ReturnValue ReturnValue
}

func (d *InitializeResponse) Type() Type {
	return TypeInitializeResponse
}

type WrappedInitializeResponse struct {
	Type    Type
	Message InitializeResponse
}

func (d *InitializeResponse) Wrap() WrappedInitializeResponse {
	return WrappedInitializeResponse{Type: d.Type(), Message: *d}
}
