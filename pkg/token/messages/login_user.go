package messages

import "github.com/edipermadi/softhsm/pkg/token/users"

type LoginUserRequest struct {
	SessionID int
	UserType  users.Type
	Pin       []byte
	Username  string
}

func (d *LoginUserRequest) Type() Type {
	return TypeLoginUserRequest
}

type WrappedLoginUserRequest struct {
	Type    Type
	Message LoginUserRequest
}

func (d *LoginUserRequest) Wrap() WrappedLoginUserRequest {
	return WrappedLoginUserRequest{Type: d.Type(), Message: *d}
}

type LoginUserResponse struct {
	ReturnValue ReturnValue
}

func (d *LoginUserResponse) Type() Type {
	return TypeLoginUserResponse
}

type WrappedLoginUserResponse struct {
	Type    Type
	Message LoginUserResponse
}

func (d *LoginUserResponse) Wrap() WrappedLoginUserResponse {
	return WrappedLoginUserResponse{Type: d.Type(), Message: *d}
}
