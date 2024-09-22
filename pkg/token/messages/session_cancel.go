package messages

type SessionCancelRequest struct {
}

func (d *SessionCancelRequest) Type() Type {
	return TypeSessionCancelRequest
}

type WrappedSessionCancelRequest struct {
	Type    Type
	Message SessionCancelRequest
}

func (d *SessionCancelRequest) Wrap() WrappedSessionCancelRequest {
	return WrappedSessionCancelRequest{Type: d.Type(), Message: *d}
}

type SessionCancelResponse struct {
	ReturnValue ReturnValue
}

func (d *SessionCancelResponse) Type() Type {
	return TypeSessionCancelResponse
}

type WrappedSessionCancelResponse struct {
	Type    Type
	Message SessionCancelResponse
}

func (d *SessionCancelResponse) Wrap() WrappedSessionCancelResponse {
	return WrappedSessionCancelResponse{Type: d.Type(), Message: *d}
}
