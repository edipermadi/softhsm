package messages

type EncryptMessageBeginRequest struct {
}

func (d *EncryptMessageBeginRequest) Type() Type {
	return TypeEncryptMessageBeginRequest
}

type WrappedEncryptMessageBeginRequest struct {
	Type    Type
	Message EncryptMessageBeginRequest
}

func (d *EncryptMessageBeginRequest) Wrap() WrappedEncryptMessageBeginRequest {
	return WrappedEncryptMessageBeginRequest{Type: d.Type(), Message: *d}
}

type EncryptMessageBeginResponse struct {
	ReturnValue ReturnValue
}

func (d *EncryptMessageBeginResponse) Type() Type {
	return TypeEncryptMessageBeginResponse
}

type WrappedEncryptMessageBeginResponse struct {
	Type    Type
	Message EncryptMessageBeginResponse
}

func (d *EncryptMessageBeginResponse) Wrap() WrappedEncryptMessageBeginResponse {
	return WrappedEncryptMessageBeginResponse{Type: d.Type(), Message: *d}
}
