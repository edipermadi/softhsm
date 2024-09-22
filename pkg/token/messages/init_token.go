package messages

type InitTokenRequest struct {
}

func (d *InitTokenRequest) Type() Type {
	return TypeInitTokenRequest
}

type WrappedInitTokenRequest struct {
	Type    Type
	Message InitTokenRequest
}

func (d *InitTokenRequest) Wrap() WrappedInitTokenRequest {
	return WrappedInitTokenRequest{Type: d.Type(), Message: *d}
}

type InitTokenResponse struct {
	ReturnValue ReturnValue
}

func (d *InitTokenResponse) Type() Type {
	return TypeInitTokenResponse
}

type WrappedInitTokenResponse struct {
	Type    Type
	Message InitTokenResponse
}

func (d *InitTokenResponse) Wrap() WrappedInitTokenResponse {
	return WrappedInitTokenResponse{Type: d.Type(), Message: *d}
}
