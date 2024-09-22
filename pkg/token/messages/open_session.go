package messages

type OpenSessionRequest struct {
}

func (d *OpenSessionRequest) Type() Type {
	return TypeOpenSessionRequest
}

type WrappedOpenSessionRequest struct {
	Type    Type
	Message OpenSessionRequest
}

func (d *OpenSessionRequest) Wrap() WrappedOpenSessionRequest {
	return WrappedOpenSessionRequest{Type: d.Type(), Message: *d}
}

type OpenSessionResponse struct {
	ReturnValue ReturnValue
}

func (d *OpenSessionResponse) Type() Type {
	return TypeOpenSessionResponse
}

type WrappedOpenSessionResponse struct {
	Type    Type
	Message OpenSessionResponse
}

func (d *OpenSessionResponse) Wrap() WrappedOpenSessionResponse {
	return WrappedOpenSessionResponse{Type: d.Type(), Message: *d}
}
