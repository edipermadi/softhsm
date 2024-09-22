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

	session, found := s.sessions[request.SessionID]
	if !found {
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_SESSION_CLOSED}, nil
	}

	if session.Context != nil {
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_SESSION_HANDLE_INVALID}, nil
	}

	mechanismType := request.Mechanism.Mechanism
	switch mechanismType {
	case messages.MechanismType_MD5:
		session.Context = &sessions.ContextDigest{
			MechanismType: mechanismType,
			Hash:          md5.New(),
		}
		s.sessions[request.SessionID] = session
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA_1:
		session.Context = &sessions.ContextDigest{
			MechanismType: mechanismType,
			Hash:          sha1.New(),
		}

		s.sessions[request.SessionID] = session
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_RIPEMD160:
		session.Context = &sessions.ContextDigest{
			MechanismType: mechanismType,
			Hash:          ripemd160.New(),
		}

		s.sessions[request.SessionID] = session
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA256:
		session.Context = &sessions.ContextDigest{
			MechanismType: mechanismType,
			Hash:          sha256.New(),
		}

		s.sessions[request.SessionID] = session
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA224:
		session.Context = &sessions.ContextDigest{
			MechanismType: mechanismType,
			Hash:          sha256.New224(),
		}

		s.sessions[request.SessionID] = session
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA384:
		session.Context = &sessions.ContextDigest{
			MechanismType: mechanismType,
			Hash:          sha512.New384(),
		}

		s.sessions[request.SessionID] = session
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA512:
		session.Context = &sessions.ContextDigest{
			MechanismType: mechanismType,
			Hash:          sha512.New(),
		}

		s.sessions[request.SessionID] = session
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA3_256:
		session.Context = &sessions.ContextDigest{
			MechanismType: mechanismType,
			Hash:          sha3.New256(),
		}

		s.sessions[request.SessionID] = session
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA3_224:
		session.Context = &sessions.ContextDigest{
			MechanismType: mechanismType,
			Hash:          sha3.New224(),
		}

		s.sessions[request.SessionID] = session
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA3_384:
		session.Context = &sessions.ContextDigest{
			MechanismType: mechanismType,
			Hash:          sha3.New384(),
		}

		s.sessions[request.SessionID] = session
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	case messages.MechanismType_SHA3_512:
		session.Context = &sessions.ContextDigest{
			MechanismType: mechanismType,
			Hash:          sha3.New512(),
		}

		s.sessions[request.SessionID] = session
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_OK}, nil

	default:
		return &messages.DigestInitResponse{ReturnValue: messages.ReturnValue_MECHANISM_INVALID}, nil
	}
}

func (s *service) Digest(request *messages.DigestRequest) (*messages.DigestResponse, error) {
	session, found := s.sessions[request.SessionID]
	if !found {
		return &messages.DigestResponse{ReturnValue: messages.ReturnValue_SESSION_CLOSED}, nil
	}

	if session.Context == nil {
		return &messages.DigestResponse{ReturnValue: messages.ReturnValue_OPERATION_NOT_INITIALIZED}, nil
	}

	if session.Context.Type() != sessions.ContextTypeDigest {
		return &messages.DigestResponse{ReturnValue: messages.ReturnValue_OPERATION_NOT_INITIALIZED}, nil
	}

	sessionDigest := session.Context.(*sessions.ContextDigest)
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

	if session.Context == nil {
		return &messages.DigestUpdateResponse{ReturnValue: messages.ReturnValue_OPERATION_NOT_INITIALIZED}, nil
	}

	if session.Context.Type() != sessions.ContextTypeDigest {
		return &messages.DigestUpdateResponse{ReturnValue: messages.ReturnValue_OPERATION_NOT_INITIALIZED}, nil
	}

	sessionDigest := session.Context.(*sessions.ContextDigest)
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

	if session.Context == nil {
		return &messages.DigestFinalResponse{ReturnValue: messages.ReturnValue_OPERATION_NOT_INITIALIZED}, nil
	}

	if session.Context.Type() != sessions.ContextTypeDigest {
		return &messages.DigestFinalResponse{ReturnValue: messages.ReturnValue_OPERATION_NOT_INITIALIZED}, nil
	}

	sessionDigest := session.Context.(*sessions.ContextDigest)
	digest := sessionDigest.Hash.Sum(nil)

	return &messages.DigestFinalResponse{ReturnValue: messages.ReturnValue_OK, Digest: digest}, nil
}
