package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/prakhar30/bankoindiana/db/sqlc"
	"github.com/prakhar30/bankoindiana/gapi"
	"github.com/prakhar30/bankoindiana/pb"
	"github.com/prakhar30/bankoindiana/utils"
	"github.com/prakhar30/bankoindiana/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGINT,
	syscall.SIGTERM,
}

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	connPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	redisOption := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOption)

	waitGroup, waitGroupContext := errgroup.WithContext(ctx)

	store := db.NewStore(connPool)
	runTaskProcessor(waitGroupContext, waitGroup, redisOption, store)
	runGPRCServer(waitGroupContext, waitGroup, config, store, taskDistributor)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error in wait group")
	}
}

func runGPRCServer(ctx context.Context, waitGroup *errgroup.Group, config utils.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create server")
	}

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listner, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	waitGroup.Go(func() error {
		log.Info().Str("port", listner.Addr().String()).Msg("started gRPC server")

		err = grpcServer.Serve(listner)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			log.Error().Err(err).Msg("gRPC server failed to serve")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("gracefully stopping gRPC server")
		grpcServer.GracefulStop()
		log.Info().Msg("gracefully stopped gRPC server")
		return nil
	})
}

func runTaskProcessor(ctx context.Context, waitGroup *errgroup.Group, redisOpt asynq.RedisClientOpt, store db.Store) {
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store)

	log.Info().Msg("starting task processor")
	err := taskProcessor.Start() // this start method already starts it in a new go routine, hence no new routine needed for it.
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("gracefully stopping task processor")
		taskProcessor.Shutdown()
		log.Info().Msg("gracefully stopped task processor")

		return nil
	})
}

// func runGinServer(config utils.Config, store db.Store) {
// 	server, err := api.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal("Cannot create server:", err)
// 	}

// 	err = server.Start(config.HTTPServerAddress)
// 	if err != nil {
// 		log.Fatal("cannot start server:", err)
// 	}
// }

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create migration")
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("cannot run migration")
	}

	log.Info().Msg("migration completed successfully")
}
