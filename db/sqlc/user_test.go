package db

import (
	"context"
	"testing"
	"time"

	"fidelis.com/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createTestUser(t *testing.T) User{
	hashedPassword,err := util.HashPassword("secret")

	require.NoError(t, err)

	arg :=CreateUserParams{
		Username: util.GenerateRandomOwner(),
		FullName: util.GenerateRandomOwner(),
		HashedPassword: hashedPassword,
		Email: util.GenerateRandomEmail(),
	}

	user, err := testQuery.CreateUser(context.Background(),arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotZero(t, user.CreatedAt)
	require.Equal(t,arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Username, user.Username)

	return user
}

func Test_CreateUser(t *testing.T) {
	createTestUser(t)
}

func Test_GetUser(t *testing.T){
	user1 := createTestUser(t)

	user2,err := testQuery.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.Equal(t,user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}