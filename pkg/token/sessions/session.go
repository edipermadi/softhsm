package sessions

import "github.com/edipermadi/softhsm/pkg/token/users"

type Session struct {
	Flags    int32
	LoggedIn bool
	UserType users.Type
	Username string
	Context  Context
}

type Context interface {
	Type() ContextType
}

type ContextType int

const (
	ContextTypeDigest = ContextType(1)
)
