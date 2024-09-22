package messages

type CloseSessionRequest struct {
	SessionID int
}

func (d *CloseSessionRequest) Type() Type {
	return TypeCloseSessionRequest
}

type WrappedCloseSessionRequest struct {
	Type    Type
	Message CloseSessionRequest
}

func (d *CloseSessionRequest) Wrap() WrappedCloseSessionRequest {
	return WrappedCloseSessionRequest{Type: d.Type(), Message: *d}
}

type CloseSessionResponse struct {
	ReturnValue ReturnValue
}

func (d *CloseSessionResponse) Type() Type {
	return TypeCloseSessionResponse
}

type WrappedCloseSessionResponse struct {
	Type    Type
	Message CloseSessionResponse
}

func (d *CloseSessionResponse) Wrap() WrappedCloseSessionResponse {
	return WrappedCloseSessionResponse{Type: d.Type(), Message: *d}
}
