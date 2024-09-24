package users

import "golang.org/x/crypto/bcrypt"

type Type int

const (
	TypeInvalid         = Type(iota)
	TypeSecurityOfficer = Type(iota)
	TypeUser            = Type(iota)
)

type Service interface {
	Authenticate(username string, password []byte) error
}

func NewService() (Service, error) {
	// FIXME add implementation
	defaultUserPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	return &service{
		defaultUserPassword: defaultUserPassword,
	}, nil
}

type service struct {
	defaultUserPassword []byte
}

func (s *service) Authenticate(username string, password []byte) error {
	// FIXME add implementation
	return bcrypt.CompareHashAndPassword(s.defaultUserPassword, password)
}
