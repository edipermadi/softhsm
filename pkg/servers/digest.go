package servers

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"

	"github.com/edipermadi/softhsm/pkg/transport/pb"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func (s *server) DigestInit(ctx context.Context, request *pb.DigestInitRequest) (*pb.DigestInitResponse, error) {
	if request.GetSessionHandle() < 1 {
		return &pb.DigestInitResponse{ReturnValue: pb.ReturnValue_SESSION_HANDLE_INVALID}, nil
	}

	if !s.sessions.Has(request.GetSessionHandle()) {
		return &pb.DigestInitResponse{ReturnValue: pb.ReturnValue_SESSION_CLOSED}, nil
	}

	if request.GetMechanism() == nil {
		return &pb.DigestInitResponse{ReturnValue: pb.ReturnValue_ARGUMENTS_BAD}, nil
	}

	session := s.sessions.Get(request.GetSessionHandle()).Value()
	switch request.GetMechanism().Mechanism {
	case pb.MechanismType_RIPEMD160:
		session.hash = ripemd160.New()
	case pb.MechanismType_MD5:
		session.hash = md5.New()
	case pb.MechanismType_SHA_1:
		session.hash = sha1.New()
	case pb.MechanismType_SHA224:
		session.hash = sha256.New224()
	case pb.MechanismType_SHA256:
		session.hash = sha256.New()
	case pb.MechanismType_SHA384:
		session.hash = sha512.New384()
	case pb.MechanismType_SHA512:
		session.hash = sha512.New()
	case pb.MechanismType_SHA3_224:
		session.hash = sha3.New224()
	case pb.MechanismType_SHA3_256:
		session.hash = sha3.New256()
	case pb.MechanismType_SHA3_384:
		session.hash = sha3.New384()
	case pb.MechanismType_SHA3_512:
		session.hash = sha3.New512()
	default:
		return &pb.DigestInitResponse{ReturnValue: pb.ReturnValue_MECHANISM_INVALID}, nil
	}

	s.sessions.Set(request.SessionHandle, session, s.ttl)
	return &pb.DigestInitResponse{ReturnValue: pb.ReturnValue_OK}, nil
}

func (s *server) Digest(ctx context.Context, request *pb.DigestRequest) (*pb.DigestResponse, error) {
	log := s.logger.WithContext(ctx)

	if request.GetSessionHandle() < 1 {
		return &pb.DigestResponse{ReturnValue: pb.ReturnValue_SESSION_HANDLE_INVALID}, nil
	}

	if !s.sessions.Has(request.GetSessionHandle()) {
		return &pb.DigestResponse{ReturnValue: pb.ReturnValue_SESSION_CLOSED}, nil
	}

	session := s.sessions.Get(request.GetSessionHandle()).Value()
	if session.hash == nil {
		return &pb.DigestResponse{ReturnValue: pb.ReturnValue_OPERATION_NOT_INITIALIZED}, nil
	}

	_, err := session.hash.Write(request.GetData())
	if err != nil {
		log.WithField("session_handle", request.GetSessionHandle()).WithError(err).Error("failed to hash data")
		return &pb.DigestResponse{ReturnValue: pb.ReturnValue_GENERAL_ERROR}, err
	}

	digest := session.hash.Sum(nil)
	return &pb.DigestResponse{ReturnValue: pb.ReturnValue_OK, Digest: digest}, nil
}

func (s *server) DigestUpdate(ctx context.Context, request *pb.DigestUpdateRequest) (*pb.DigestUpdateResponse, error) {
	log := s.logger.WithContext(ctx)

	if request.GetSessionHandle() < 1 {
		return &pb.DigestUpdateResponse{ReturnValue: pb.ReturnValue_SESSION_HANDLE_INVALID}, nil
	}

	if !s.sessions.Has(request.GetSessionHandle()) {
		return &pb.DigestUpdateResponse{ReturnValue: pb.ReturnValue_SESSION_CLOSED}, nil
	}

	session := s.sessions.Get(request.GetSessionHandle()).Value()
	if session.hash == nil {
		return &pb.DigestUpdateResponse{ReturnValue: pb.ReturnValue_OPERATION_NOT_INITIALIZED}, nil
	}

	_, err := session.hash.Write(request.GetData())
	if err != nil {
		log.WithField("session_handle", request.GetSessionHandle()).WithError(err).Error("failed to hash data")
		return &pb.DigestUpdateResponse{ReturnValue: pb.ReturnValue_GENERAL_ERROR}, err
	}

	s.sessions.Set(request.GetSessionHandle(), session, s.ttl)
	return &pb.DigestUpdateResponse{ReturnValue: pb.ReturnValue_OK}, nil
}

func (s *server) DigestKey(ctx context.Context, request *pb.DigestKeyRequest) (*pb.DigestKeyResponse, error) {
	// TODO add implementation
	return nil, errors.New("not implemented")
}

func (s *server) DigestFinal(ctx context.Context, request *pb.DigestFinalRequest) (*pb.DigestFinalResponse, error) {
	if request.GetSessionHandle() < 1 {
		return &pb.DigestFinalResponse{ReturnValue: pb.ReturnValue_SESSION_HANDLE_INVALID}, nil
	}

	if !s.sessions.Has(request.GetSessionHandle()) {
		return &pb.DigestFinalResponse{ReturnValue: pb.ReturnValue_SESSION_CLOSED}, nil
	}

	session := s.sessions.Get(request.GetSessionHandle()).Value()
	if session.hash == nil {
		return &pb.DigestFinalResponse{ReturnValue: pb.ReturnValue_OPERATION_NOT_INITIALIZED}, nil
	}

	digest := session.hash.Sum(nil)
	return &pb.DigestFinalResponse{ReturnValue: pb.ReturnValue_OK, Digest: digest}, nil
}
