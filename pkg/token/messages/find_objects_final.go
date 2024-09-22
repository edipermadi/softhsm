package messages

type FindObjectsFinalRequest struct {
}

func (d *FindObjectsFinalRequest) Type() Type {
	return TypeFindObjectsFinalRequest
}

type WrappedFindObjectsFinalRequest struct {
	Type    Type
	Message FindObjectsFinalRequest
}

func (d *FindObjectsFinalRequest) Wrap() WrappedFindObjectsFinalRequest {
	return WrappedFindObjectsFinalRequest{Type: d.Type(), Message: *d}
}

type FindObjectsFinalResponse struct {
	ReturnValue ReturnValue
}

func (d *FindObjectsFinalResponse) Type() Type {
	return TypeFindObjectsFinalResponse
}

type WrappedFindObjectsFinalResponse struct {
	Type    Type
	Message FindObjectsFinalResponse
}

func (d *FindObjectsFinalResponse) Wrap() WrappedFindObjectsFinalResponse {
	return WrappedFindObjectsFinalResponse{Type: d.Type(), Message: *d}
}
