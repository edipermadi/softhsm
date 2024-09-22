package messages

type FinalizeRequest struct {
}

func (d *FinalizeRequest) Type() Type {
	return TypeFinalizeRequest
}

type WrappedFinalizeRequest struct {
	Type    Type
	Message FinalizeRequest
}

func (d *FinalizeRequest) Wrap() WrappedFinalizeRequest {
	return WrappedFinalizeRequest{Type: d.Type(), Message: *d}
}

type FinalizeResponse struct {
	ReturnValue ReturnValue
}

func (d *FinalizeResponse) Type() Type {
	return TypeFinalizeResponse
}

type WrappedFinalizeResponse struct {
	Type    Type
	Message FinalizeResponse
}

func (d *FinalizeResponse) Wrap() WrappedFinalizeResponse {
	return WrappedFinalizeResponse{Type: d.Type(), Message: *d}
}
