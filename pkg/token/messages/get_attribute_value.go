package messages

type GetAttributeValueRequest struct {
}

func (d *GetAttributeValueRequest) Type() Type {
	return TypeGetAttributeValueRequest
}

type WrappedGetAttributeValueRequest struct {
	Type    Type
	Message GetAttributeValueRequest
}

func (d *GetAttributeValueRequest) Wrap() WrappedGetAttributeValueRequest {
	return WrappedGetAttributeValueRequest{Type: d.Type(), Message: *d}
}

type GetAttributeValueResponse struct {
	ReturnValue ReturnValue
}

func (d *GetAttributeValueResponse) Type() Type {
	return TypeGetAttributeValueResponse
}

type WrappedGetAttributeValueResponse struct {
	Type    Type
	Message GetAttributeValueResponse
}

func (d *GetAttributeValueResponse) Wrap() WrappedGetAttributeValueResponse {
	return WrappedGetAttributeValueResponse{Type: d.Type(), Message: *d}
}
