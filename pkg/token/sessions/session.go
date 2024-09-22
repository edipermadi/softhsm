package sessions

type Session struct {
	Flags   int32
	Context Context
}
type Context interface {
	Type() ContextType
}

type ContextType int

const (
	ContextTypeDigest = ContextType(1)
)
