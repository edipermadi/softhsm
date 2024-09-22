package messages

type FindObjectsRequest struct {
}

func (d *FindObjectsRequest) Type() Type {
	return TypeFindObjectsRequest
}

type WrappedFindObjectsRequest struct {
	Type    Type
	Message FindObjectsRequest
}

func (d *FindObjectsRequest) Wrap() WrappedFindObjectsRequest {
	return WrappedFindObjectsRequest{Type: d.Type(), Message: *d}
}

type FindObjectsResponse struct {
	ReturnValue ReturnValue
}

func (d *FindObjectsResponse) Type() Type {
	return TypeFindObjectsResponse
}

type WrappedFindObjectsResponse struct {
	Type    Type
	Message FindObjectsResponse
}

func (d *FindObjectsResponse) Wrap() WrappedFindObjectsResponse {
	return WrappedFindObjectsResponse{Type: d.Type(), Message: *d}
}
