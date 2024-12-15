package main

import (
	"context"
	"log/slog"
	"os"

	"connectrpc.com/connect"
	"connectrpc.com/validate"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jchapman63/neo-gkmn/internal/config"
	"github.com/jchapman63/neo-gkmn/internal/database"
	"github.com/jchapman63/neo-gkmn/internal/server"
	"github.com/jchapman63/neo-gkmn/internal/service/gkmn"
)

const ClientCMD = "client"
const MigrationCMD = "migrate"

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("failure to load config", "err", err)
		os.Exit(1)
	}

	pool, err := pgxpool.New(ctx, cfg.Database.URL)
	if err != nil {
		slog.Error("failure to create db connection", "err", err)
	}

	api, err := server.New(&cfg.Server)
	if err != nil {
		slog.Error("failure to create server", "err", err)
		os.Exit(1)
	}

	validator, err := validate.NewInterceptor()
	if err != nil {
		slog.Error("failure to create interceptor", "err", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		switch os.Args[0] {
		case ClientCMD:
		case MigrationCMD:
		}
	}

	gameService := gkmn.NewGameService(
		database.New(pool),
		gkmn.WithHandlerOptions(
			connect.WithInterceptors(validator),
		),
	)

	api.RegisterService(gameService)

	if err := api.Serve(); err != nil {
		slog.Error("failed to serve application", "error", err)
		os.Exit(1)
	}
}
