package messages

type SetAttributeValueRequest struct {
}

func (d *SetAttributeValueRequest) Type() Type {
	return TypeSetAttributeValueRequest
}

type WrappedSetAttributeValueRequest struct {
	Type    Type
	Message SetAttributeValueRequest
}

func (d *SetAttributeValueRequest) Wrap() WrappedSetAttributeValueRequest {
	return WrappedSetAttributeValueRequest{Type: d.Type(), Message: *d}
}

type SetAttributeValueResponse struct {
	ReturnValue ReturnValue
}

func (d *SetAttributeValueResponse) Type() Type {
	return TypeSetAttributeValueResponse
}

type WrappedSetAttributeValueResponse struct {
	Type    Type
	Message SetAttributeValueResponse
}

func (d *SetAttributeValueResponse) Wrap() WrappedSetAttributeValueResponse {
	return WrappedSetAttributeValueResponse{Type: d.Type(), Message: *d}
}
