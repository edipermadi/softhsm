package services_test

import (
	"testing"

	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/edipermadi/softhsm/pkg/token/services"
	"github.com/edipermadi/softhsm/pkg/token/users"
	"github.com/stretchr/testify/require"
)

func TestService_OpenSession(t *testing.T) {
	userService, err := users.NewService()
	require.NoError(t, err)

	svc := services.NewService(userService)
	openSessionResponse, err := svc.OpenSession(&messages.OpenSessionRequest{Flags: 0})
	require.NoError(t, err)
	require.Equal(t, messages.ReturnValue_OK, openSessionResponse.ReturnValue)
	require.NotEmpty(t, openSessionResponse.SessionID)
}
