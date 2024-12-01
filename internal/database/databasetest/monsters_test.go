package databasetest

import (
	"context"

	"github.com/jchapman63/neo-gkmn/internal/database"
)

// tests Querier monsters methods with actual db queries

func (s *DatabaseTestSuite) Test_ListMonsters() {
	ctx := context.Background()
	db := database.New(s.tx)

	monster := database.CreateMonsterParams{
		ID:     "11111111-1111-1111-1111-111111111111",
		Name:   "Bulbasaur",
		Type:   "Grass",
		Basehp: 50,
	}
	err := db.CreateMonster(ctx, monster)
	s.NoError(err)

	monList, err := db.ListMonsters(ctx)
	s.NoError(err)

	s.Equal(len(monList), 1)
	s.Equal(monList[0].ID, monster.ID)

}

func (s *DatabaseTestSuite) Test_CreateAndFetchMonster() {
	ctx := context.Background()
	db := database.New(s.tx)

	monster := database.CreateMonsterParams{
		ID:     "11111111-1111-1111-1111-111111111111",
		Name:   "Bulbasaur",
		Type:   "Grass",
		Basehp: 50,
	}
	err := db.CreateMonster(ctx, monster)
	s.NoError(err)

	dbMon, err := db.FetchMonster(ctx, monster.ID)
	s.NoError(err)

	s.Equal(monster.ID, dbMon.ID)
}

func (s *DatabaseTestSuite) Test_CreateAndFetchMove() {
	ctx := context.Background()
	db := database.New(s.tx)

	move := database.CreateMoveParams{
		ID:    "11111111-1111-1111-1111-111111111111",
		Name:  "tackle",
		Power: 10,
		Type:  "Normal",
	}
	err := db.CreateMove(ctx, move)
	s.NoError(err)

	dbMove, err := db.FetchMove(ctx, move.ID)
	s.NoError(err)

	s.Equal(move.ID, dbMove.ID)
}
