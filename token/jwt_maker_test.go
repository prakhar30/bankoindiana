package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/prakhar30/bankoindiana/utils"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	randomString := utils.RandomString(32)
	t.Log(randomString)
	maker, err := NewJWTMaker(randomString)
	require.NoError(t, err)
	t.Log(maker)

	username := utils.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	randomString := utils.RandomString(32)
	maker, err := NewJWTMaker(randomString)
	require.NoError(t, err)

	token, payload, err := maker.CreateToken(utils.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgoNone(t *testing.T) {
	payload, err := NewPayload(utils.RandomOwner(), time.Minute)
	require.NoError(t, err)
	wrappedPayload := NewWrappedPayload(payload)
	require.NotNil(t, wrappedPayload)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, wrappedPayload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	randomString := utils.RandomString(32)
	maker, err := NewJWTMaker(randomString)
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
}
