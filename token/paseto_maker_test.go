package token

import (
	"testing"
	"time"

	"github.com/prakhar30/bankoindiana/utils"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	randomString := utils.RandomString(32)
	t.Log(randomString)
	maker, err := NewPasetoMaker(randomString)
	require.NoError(t, err)
	t.Log(maker)

	username := utils.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)
	role := utils.DepositorRole

	token, payload, err := maker.CreateToken(username, role, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.Equal(t, role, payload.Role)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	randomString := utils.RandomString(32)
	maker, err := NewPasetoMaker(randomString)
	require.NoError(t, err)
	token, payload, err := maker.CreateToken(utils.RandomOwner(), utils.DepositorRole, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
}
