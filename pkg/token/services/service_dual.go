package services

import (
	"errors"
	"github.com/edipermadi/softhsm/pkg/token/messages"
)

func (s *service) DigestEncryptUpdate(request *messages.DigestEncryptUpdateRequest) (*messages.DigestEncryptUpdateResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) DecryptDigestUpdate(request *messages.DecryptDigestUpdateRequest) (*messages.DecryptDigestUpdateResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) SignEncryptUpdate(request *messages.SignEncryptUpdateRequest) (*messages.SignEncryptUpdateResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) DecryptVerifyUpdate(request *messages.DecryptVerifyUpdateRequest) (*messages.DecryptVerifyUpdateResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}
