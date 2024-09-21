package messages

type Type int

const (
	TypeInvalid              = Type(iota)
	TypeDigestInitRequest    = Type(iota)
	TypeDigestInitResponse   = Type(iota)
	TypeDigestRequest        = Type(iota)
	TypeDigestResponse       = Type(iota)
	TypeDigestUpdateRequest  = Type(iota)
	TypeDigestUpdateResponse = Type(iota)
	TypeDigestFinalRequest   = Type(iota)
	TypeDigestFinalResponse  = Type(iota)
)

type Message interface {
	Type() Type
}
