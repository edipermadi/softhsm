package servers

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

	var response messages.Message
	switch message.Type() {
	case messages.TypeDigestInitRequest:
		response, err = s.service.DigestInit(message.(*messages.DigestInitRequest))
	case messages.TypeDigestRequest:
		response, err = s.service.Digest(message.(*messages.DigestRequest))
	case messages.TypeDigestUpdateRequest:
		response, err = s.service.DigestUpdate(message.(*messages.DigestUpdateRequest))
	case messages.TypeDigestFinalRequest:
		response, err = s.service.DigestFinal(message.(*messages.DigestFinalRequest))
	default:
		return nil, errors.New("unsupported message")
	}

	if err != nil {
		return nil, err
	}

	return messages.Marshall(response)
}
