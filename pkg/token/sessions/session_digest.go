package sessions

import (
	"hash"

	"github.com/edipermadi/softhsm/pkg/token/messages"
)

type SessionDigest struct {
	MechanismType messages.MechanismType
	Hash          hash.Hash
	DigestSize    int // digest size
}

func (s *SessionDigest) Type() SessionType {
	return SessionTypeDigest
}
