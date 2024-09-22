package messages

type VerifyInitRequest struct {
}

func (d *VerifyInitRequest) Type() Type {
	return TypeVerifyInitRequest
}

type WrappedVerifyInitRequest struct {
	Type    Type
	Message VerifyInitRequest
}

func (d *VerifyInitRequest) Wrap() WrappedVerifyInitRequest {
	return WrappedVerifyInitRequest{Type: d.Type(), Message: *d}
}

type VerifyInitResponse struct {
	ReturnValue ReturnValue
}

func (d *VerifyInitResponse) Type() Type {
	return TypeVerifyInitResponse
}

type WrappedVerifyInitResponse struct {
	Type    Type
	Message VerifyInitResponse
}

func (d *VerifyInitResponse) Wrap() WrappedVerifyInitResponse {
	return WrappedVerifyInitResponse{Type: d.Type(), Message: *d}
}
