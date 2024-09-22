package messages

type InitPINRequest struct {
}

func (d *InitPINRequest) Type() Type {
	return TypeInitPINRequest
}

type WrappedInitPINRequest struct {
	Type    Type
	Message InitPINRequest
}

func (d *InitPINRequest) Wrap() WrappedInitPINRequest {
	return WrappedInitPINRequest{Type: d.Type(), Message: *d}
}

type InitPINResponse struct {
	ReturnValue ReturnValue
}

func (d *InitPINResponse) Type() Type {
	return TypeInitPINResponse
}

type WrappedInitPINResponse struct {
	Type    Type
	Message InitPINResponse
}

func (d *InitPINResponse) Wrap() WrappedInitPINResponse {
	return WrappedInitPINResponse{Type: d.Type(), Message: *d}
}
