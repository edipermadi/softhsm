package messages

type DecryptMessageBeginRequest struct {
}

func (d *DecryptMessageBeginRequest) Type() Type {
	return TypeDecryptMessageBeginRequest
}

type WrappedDecryptMessageBeginRequest struct {
	Type    Type
	Message DecryptMessageBeginRequest
}

func (d *DecryptMessageBeginRequest) Wrap() WrappedDecryptMessageBeginRequest {
	return WrappedDecryptMessageBeginRequest{Type: d.Type(), Message: *d}
}

type DecryptMessageBeginResponse struct {
	ReturnValue ReturnValue
}

func (d *DecryptMessageBeginResponse) Type() Type {
	return TypeDecryptMessageBeginResponse
}

type WrappedDecryptMessageBeginResponse struct {
	Type    Type
	Message DecryptMessageBeginResponse
}

func (d *DecryptMessageBeginResponse) Wrap() WrappedDecryptMessageBeginResponse {
	return WrappedDecryptMessageBeginResponse{Type: d.Type(), Message: *d}
}
