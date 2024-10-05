package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jchapman63/neo-gkmn/internal/database"
	"github.com/jchapman63/neo-gkmn/internal/pkg"
)

func main() {
	// quick goal
	testBattle()
}

func testBattle() {
	// simulate db queried bulbasaur
	bulbasaur:= database.Monster{
		ID: uuid.New(),
		Name: "bulbasaur",
		Type: "grass",
		Basehp: 32,
	}

	// move from db
	move := database.Move{
		ID: uuid.New(),
		Name: "tackle",
		Power: 10,
	}
	// server will spawn a battle bulbasaur from db bulbasaur
	battleBulba := pkg.BattleMon{
		Monster: &bulbasaur,
		LiveHp: bulbasaur.Basehp,
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
