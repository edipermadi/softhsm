package messages

type SeedRandomRequest struct {
}

func (d *SeedRandomRequest) Type() Type {
	return TypeSeedRandomRequest
}

type WrappedSeedRandomRequest struct {
	Type    Type
	Message SeedRandomRequest
}

func (d *SeedRandomRequest) Wrap() WrappedSeedRandomRequest {
	return WrappedSeedRandomRequest{Type: d.Type(), Message: *d}
}

type SeedRandomResponse struct {
	ReturnValue ReturnValue
}

func (d *SeedRandomResponse) Type() Type {
	return TypeSeedRandomResponse
}

type WrappedSeedRandomResponse struct {
	Type    Type
	Message SeedRandomResponse
}

func (d *SeedRandomResponse) Wrap() WrappedSeedRandomResponse {
	return WrappedSeedRandomResponse{Type: d.Type(), Message: *d}
}
