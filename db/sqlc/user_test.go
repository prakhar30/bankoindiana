package db

import (
	"context"
	"testing"
	"time"

	"github.com/prakhar30/bankoindiana/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(password string) (CreateUserParams, User, error) {
	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: password,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}
	user, err := testStore.CreateUser(context.Background(), arg)
	return arg, user, err
}

func TestCreateUser(t *testing.T) {
	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)

	arg, user, err := createRandomUser(hashedPassword)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.Time.IsZero())
}

func TestGetUser(t *testing.T) {
	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)

	_, user, _ := createRandomUser(hashedPassword)

	fetchedUser, err := testStore.GetUser(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, fetchedUser)

	require.Equal(t, user.Username, fetchedUser.Username)
	require.Equal(t, user.HashedPassword, fetchedUser.HashedPassword)
	require.Equal(t, user.Email, fetchedUser.Email)
	require.Equal(t, user.FullName, fetchedUser.FullName)
	require.WithinDuration(t, user.CreatedAt.Time, fetchedUser.CreatedAt.Time, time.Second)
}
