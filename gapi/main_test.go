package gapi

import (
	"testing"
	"time"

	db "github.com/prakhar30/bankoindiana/db/sqlc"
	"github.com/prakhar30/bankoindiana/utils"
	"github.com/prakhar30/bankoindiana/worker"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}
