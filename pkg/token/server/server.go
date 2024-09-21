package server

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/edipermadi/softhsm/pkg/token/services"
)

type Server interface {
	Process(data []byte) ([]byte, error)
}

func NewServer(service services.Service) Server {
	return &server{service: service}
}

type server struct {
	service services.Service
}

func (s *server) Process(data []byte) ([]byte, error) {
	message, err := messages.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	switch message.Type() {
	case messages.TypeDigestInitRequest:
	case messages.TypeDigestRequest:
	case messages.TypeDigestUpdateRequest:
	case messages.TypeDigestFinalRequest:
	default:
		return nil, errors.New("unsupported message")
	}
}
