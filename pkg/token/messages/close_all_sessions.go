package messages

type CloseAllSessionsRequest struct {
}

func (d *CloseAllSessionsRequest) Type() Type {
	return TypeCloseAllSessionsRequest
}

type WrappedCloseAllSessionsRequest struct {
	Type    Type
	Message CloseAllSessionsRequest
}

func (d *CloseAllSessionsRequest) Wrap() WrappedCloseAllSessionsRequest {
	return WrappedCloseAllSessionsRequest{Type: d.Type(), Message: *d}
}

type CloseAllSessionsResponse struct {
	ReturnValue ReturnValue
}

func (d *CloseAllSessionsResponse) Type() Type {
	return TypeCloseAllSessionsResponse
}

type WrappedCloseAllSessionsResponse struct {
	Type    Type
	Message CloseAllSessionsResponse
}

func (d *CloseAllSessionsResponse) Wrap() WrappedCloseAllSessionsResponse {
	return WrappedCloseAllSessionsResponse{Type: d.Type(), Message: *d}
}
