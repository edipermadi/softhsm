package messages

type LogoutRequest struct {
	SessionID int
}

func (d *LogoutRequest) Type() Type {
	return TypeLogoutRequest
}

type WrappedLogoutRequest struct {
	Type    Type
	Message LogoutRequest
}

func (d *LogoutRequest) Wrap() WrappedLogoutRequest {
	return WrappedLogoutRequest{Type: d.Type(), Message: *d}
}

type LogoutResponse struct {
	ReturnValue ReturnValue
}

func (d *LogoutResponse) Type() Type {
	return TypeLogoutResponse
}

type WrappedLogoutResponse struct {
	Type    Type
	Message LogoutResponse
}

func (d *LogoutResponse) Wrap() WrappedLogoutResponse {
	return WrappedLogoutResponse{Type: d.Type(), Message: *d}
}
