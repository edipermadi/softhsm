package sessions

import (
	"hash"

	"github.com/edipermadi/softhsm/pkg/token/messages"
)

type ContextDigest struct {
	MechanismType messages.MechanismType
	Hash          hash.Hash
	DigestSize    int // digest size
}

func (s *ContextDigest) Type() ContextType {
	return ContextTypeDigest
}
