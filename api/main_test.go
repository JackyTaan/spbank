package api

import (
	"os"
	"testing"
	"time"

	db "github.com/JackyTaan/spbank/db/sqlc"
	"github.com/JackyTaan/spbank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	Server, err := NewServer(config, store)
	require.NoError(t, err)

	return Server
}

func TestMain(m *testing.M) {
	//
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
	//
}
