package messages

type FindObjectsInitRequest struct {
}

func (d *FindObjectsInitRequest) Type() Type {
	return TypeFindObjectsInitRequest
}

type WrappedFindObjectsInitRequest struct {
	Type    Type
	Message FindObjectsInitRequest
}

func (d *FindObjectsInitRequest) Wrap() WrappedFindObjectsInitRequest {
	return WrappedFindObjectsInitRequest{Type: d.Type(), Message: *d}
}

type FindObjectsInitResponse struct {
	ReturnValue ReturnValue
}

func (d *FindObjectsInitResponse) Type() Type {
	return TypeFindObjectsInitResponse
}

type WrappedFindObjectsInitResponse struct {
	Type    Type
	Message FindObjectsInitResponse
}

func (d *FindObjectsInitResponse) Wrap() WrappedFindObjectsInitResponse {
	return WrappedFindObjectsInitResponse{Type: d.Type(), Message: *d}
}
