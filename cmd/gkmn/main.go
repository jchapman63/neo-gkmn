package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/jchapman63/neo-gkmn/internal/config"
	"github.com/jchapman63/neo-gkmn/internal/database"
	"github.com/jchapman63/neo-gkmn/internal/pkg"
	"github.com/jchapman63/neo-gkmn/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("failure to load config", "err", err)
		os.Exit(1)
	}

	api, err := server.New(&cfg.Server)
	if err != nil {
		slog.Error("failure to create server", "err", err)
	}

	api.Serve()
}

// TODO - remove
func testBattle() {
	// simulate db queried bulbasaur
	bulbasaur := database.Monster{
		ID:     uuid.New(),
		Name:   "bulbasaur",
		Type:   "grass",
		Basehp: 32,
	}

	// move from db
	move := database.Move{
		ID:    uuid.New(),
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
