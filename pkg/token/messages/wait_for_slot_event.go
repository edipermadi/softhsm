package messages

type WaitForSlotEventRequest struct {
}

func (d *WaitForSlotEventRequest) Type() Type {
	return TypeWaitForSlotEventRequest
}

type WrappedWaitForSlotEventRequest struct {
	Type    Type
	Message WaitForSlotEventRequest
}

func (d *WaitForSlotEventRequest) Wrap() WrappedWaitForSlotEventRequest {
	return WrappedWaitForSlotEventRequest{Type: d.Type(), Message: *d}
}

type WaitForSlotEventResponse struct {
	ReturnValue ReturnValue
}

func (d *WaitForSlotEventResponse) Type() Type {
	return TypeWaitForSlotEventResponse
}

type WrappedWaitForSlotEventResponse struct {
	Type    Type
	Message WaitForSlotEventResponse
}

func (d *WaitForSlotEventResponse) Wrap() WrappedWaitForSlotEventResponse {
	return WrappedWaitForSlotEventResponse{Type: d.Type(), Message: *d}
}
