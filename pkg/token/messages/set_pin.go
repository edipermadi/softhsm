package messages

type SetPINRequest struct {
}

func (d *SetPINRequest) Type() Type {
	return TypeSetPINRequest
}

type WrappedSetPINRequest struct {
	Type    Type
	Message SetPINRequest
}

func (d *SetPINRequest) Wrap() WrappedSetPINRequest {
	return WrappedSetPINRequest{Type: d.Type(), Message: *d}
}

type SetPINResponse struct {
	ReturnValue ReturnValue
}

func (d *SetPINResponse) Type() Type {
	return TypeSetPINResponse
}

type WrappedSetPINResponse struct {
	Type    Type
	Message SetPINResponse
}

func (d *SetPINResponse) Wrap() WrappedSetPINResponse {
	return WrappedSetPINResponse{Type: d.Type(), Message: *d}
}
