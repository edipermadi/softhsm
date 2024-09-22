package services

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"

	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/edipermadi/softhsm/pkg/token/sessions"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func (s *service) DigestInit(request *messages.DigestInitRequest) (*messages.DigestInitResponse, error) {
	if request == nil {
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_ARGUMENTS_BAD}, nil
	}

	mechanismType := request.Mechanism.Mechanism
	switch mechanismType {
	case messages.MechanismType_MD5:
		s.sessions[request.SessionID] = &sessions.SessionDigest{
			MechanismType: mechanismType,
			Hash:          md5.New(),
		}

	case messages.MechanismType_SHA_1:
		s.sessions[request.SessionID] = &sessions.SessionDigest{
			MechanismType: mechanismType,
			Hash:          sha1.New(),
		}
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_RIPEMD160:
		s.sessions[request.SessionID] = &sessions.SessionDigest{
			MechanismType: mechanismType,
			Hash:          ripemd160.New(),
		}
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA256:
		s.sessions[request.SessionID] = &sessions.SessionDigest{
			MechanismType: mechanismType,
			Hash:          sha256.New(),
		}
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA224:
		s.sessions[request.SessionID] = &sessions.SessionDigest{
			MechanismType: mechanismType,
			Hash:          sha256.New224(),
		}
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA384:
		s.sessions[request.SessionID] = &sessions.SessionDigest{
			MechanismType: mechanismType,
			Hash:          sha512.New384(),
		}
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA512:
		s.sessions[request.SessionID] = &sessions.SessionDigest{
			MechanismType: mechanismType,
			Hash:          sha512.New(),
		}
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA3_256:
		s.sessions[request.SessionID] = &sessions.SessionDigest{
			MechanismType: mechanismType,
			Hash:          sha3.New256(),
		}
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA3_224:
		s.sessions[request.SessionID] = &sessions.SessionDigest{
			MechanismType: mechanismType,
			Hash:          sha3.New224(),
		}
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA3_384:
		s.sessions[request.SessionID] = &sessions.SessionDigest{
			MechanismType: mechanismType,
			Hash:          sha3.New384(),
		}
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA3_512:
		s.sessions[request.SessionID] = &sessions.SessionDigest{
			MechanismType: mechanismType,
			Hash:          sha3.New512(),
		}
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	default:
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_MECHANISM_INVALID}, nil
	}

	return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil
}

func (s *service) Digest(request *messages.DigestRequest) (*messages.DigestResponse, error) {
	session, found := s.sessions[request.SessionID]
	if !found {
		return &messages.DigestResponse{ReturnValue: messages.ReturnValue_SESSION_CLOSED}, nil
	}

	if session.Type() != sessions.SessionTypeDigest {
		return &messages.DigestResponse{ReturnValue: messages.ReturnValue_OPERATION_NOT_INITIALIZED}, nil
	}

	sessionDigest := session.(*sessions.SessionDigest)
	_, err := sessionDigest.Hash.Write(request.Data)
	if err != nil {
		return nil, err
	}

	digest := sessionDigest.Hash.Sum(nil)
	return &messages.DigestResponse{ReturnValue: messages.ReturnValue_OK, Digest: digest}, nil
}

func (s *service) DigestUpdate(request *messages.DigestUpdateRequest) (*messages.DigestUpdateResponse, error) {
	session, found := s.sessions[request.SessionID]
	if !found {
		return &messages.DigestUpdateResponse{ReturnValue: messages.ReturnValue_SESSION_CLOSED}, nil
	}

	if session.Type() != sessions.SessionTypeDigest {
		return &messages.DigestUpdateResponse{ReturnValue: messages.ReturnValue_OPERATION_NOT_INITIALIZED}, nil
	}

	sessionDigest := session.(*sessions.SessionDigest)
	_, err := sessionDigest.Hash.Write(request.Data)
	if err != nil {
		return nil, err
	}

	return &messages.DigestUpdateResponse{ReturnValue: messages.ReturnValue_OK}, nil
}

func (s *service) DigestKey(request *messages.DigestKeyRequest) (*messages.DigestKeyResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *service) DigestFinal(request *messages.DigestFinalRequest) (*messages.DigestFinalResponse, error) {
	session, found := s.sessions[request.SessionID]
	if !found {
		return &messages.DigestFinalResponse{ReturnValue: messages.ReturnValue_SESSION_CLOSED}, nil
	}

	if session.Type() != sessions.SessionTypeDigest {
		return &messages.DigestFinalResponse{ReturnValue: messages.ReturnValue_OPERATION_NOT_INITIALIZED}, nil
	}

	sessionDigest := session.(*sessions.SessionDigest)
	digest := sessionDigest.Hash.Sum(nil)

	return &messages.DigestFinalResponse{ReturnValue: messages.ReturnValue_OK, Digest: digest}, nil
}
