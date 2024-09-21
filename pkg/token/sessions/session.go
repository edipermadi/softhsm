package sessions

type Session interface {
	Type() SessionType
}

type SessionType int

const (
	SessionTypeDigest = SessionType(1)
)
