package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"connectrpc.com/connect"
	"connectrpc.com/validate"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jchapman63/neo-gkmn/internal/config"
	"github.com/jchapman63/neo-gkmn/internal/database"
	"github.com/jchapman63/neo-gkmn/internal/pkg"
	"github.com/jchapman63/neo-gkmn/internal/server"
	"github.com/jchapman63/neo-gkmn/internal/service/gkmn"
)

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

	gameService := gkmn.NewGameService(
		database.New(pool),
		gkmn.WithHandlerOptions(
			connect.WithInterceptors(validator),
		),
	)

	api.RegisterService(gameService)
	go gameService.Listen(&ctx)

	api.Serve()
}

// TODO - remove after I seed the database
func testBattle() {
	// simulate db queried bulbasaur
	bulbasaur := database.Monster{
		ID:     uuid.NewString(),
		Name:   "bulbasaur",
		Type:   "grass",
		Basehp: 32,
	}

	// move from db
	move := database.Move{
		ID:    uuid.NewString(),
		Name:  "tackle",
		Power: 10,
	}
	// server will spawn a battle bulbasaur from db bulbasaur
	battleBulba := pkg.BattleMon{
		Monster: &bulbasaur,
		LiveHp:  bulbasaur.Basehp,
	}

	battle := pkg.Battle{
		Monsters: []*pkg.BattleMon{
			&battleBulba,
		},
	}

	fmt.Println("before battle bulba", battleBulba.LiveHp)
	battle.Damage(bulbasaur.ID, move)
	fmt.Println("after battle bulba", battleBulba.LiveHp)
}
