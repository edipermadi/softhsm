package services

import (
	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/edipermadi/softhsm/pkg/token/sessions"
)

type Service interface {
	DigestInit(request *messages.DigestInitRequest) (*messages.DigestInitResponse, error)
	Digest(request *messages.DigestRequest) (*messages.DigestResponse, error)
	DigestUpdate(request *messages.DigestUpdateRequest) (*messages.DigestUpdateResponse, error)
	DigestFinal(request *messages.DigestFinalRequest) (*messages.DigestFinalResponse, error)
}

func NewService() Service {
	return &service{sessions: make(map[int]sessions.Session)}
}

type service struct {
	sessions map[int]sessions.Session
}
