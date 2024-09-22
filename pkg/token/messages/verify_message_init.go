package messages

type VerifyMessageInitRequest struct {
}

func (d *VerifyMessageInitRequest) Type() Type {
	return TypeVerifyMessageInitRequest
}

type WrappedVerifyMessageInitRequest struct {
	Type    Type
	Message VerifyMessageInitRequest
}

func (d *VerifyMessageInitRequest) Wrap() WrappedVerifyMessageInitRequest {
	return WrappedVerifyMessageInitRequest{Type: d.Type(), Message: *d}
}

type VerifyMessageInitResponse struct {
	ReturnValue ReturnValue
}

func (d *VerifyMessageInitResponse) Type() Type {
	return TypeVerifyMessageInitResponse
}

type WrappedVerifyMessageInitResponse struct {
	Type    Type
	Message VerifyMessageInitResponse
}

func (d *VerifyMessageInitResponse) Wrap() WrappedVerifyMessageInitResponse {
	return WrappedVerifyMessageInitResponse{Type: d.Type(), Message: *d}
}
