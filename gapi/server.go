package gapi

import (
	"fmt"

	db "github.com/prakhar30/bankoindiana/db/sqlc"
	"github.com/prakhar30/bankoindiana/pb"
	"github.com/prakhar30/bankoindiana/token"
	"github.com/prakhar30/bankoindiana/utils"
	"github.com/prakhar30/bankoindiana/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config          utils.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config utils.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
