package messages

type LoginRequest struct {
}

func (d *LoginRequest) Type() Type {
	return TypeLoginRequest
}

type WrappedLoginRequest struct {
	Type    Type
	Message LoginRequest
}

func (d *LoginRequest) Wrap() WrappedLoginRequest {
	return WrappedLoginRequest{Type: d.Type(), Message: *d}
}

type LoginResponse struct {
	ReturnValue ReturnValue
}

func (d *LoginResponse) Type() Type {
	return TypeLoginResponse
}

type WrappedLoginResponse struct {
	Type    Type
	Message LoginResponse
}

func (d *LoginResponse) Wrap() WrappedLoginResponse {
	return WrappedLoginResponse{Type: d.Type(), Message: *d}
}
