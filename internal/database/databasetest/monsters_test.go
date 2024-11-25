package databasetest

import (
	"context"

	"github.com/jchapman63/neo-gkmn/internal/database"
)

func NewDatabaseMonsterParams() database.CreateMonsterParams {
	return database.CreateMonsterParams{
		ID:     "11111111-1111-1111-1111-111111111111",
		Name:   "Bulbasaur",
		Type:   "Grass",
		Basehp: 50,
	}
}

func NewDatabaseMoveParams() database.CreateMoveParams {
	return database.CreateMoveParams{
		ID:    "11111111-1111-1111-1111-111111111111",
		Name:  "tackle",
		Power: 10,
		Type:  "Normal",
	}
}

func NewDatabaseStatParams(monsterID string) database.CreateStatForMonParams {
	return database.CreateStatForMonParams{
		Monsterid: monsterID,
		Stattype:  "speed",
		Power:     40,
	}
}

func NewDatabaseFetchStatParams(monsterID string, statType string) database.FetchStatParams {
	return database.FetchStatParams{
		MonsterID: monsterID,
		StatType:  statType,
	}
}

func NewRegisterMoveParams(monsterID string, moveID string) database.RegisterMoveForMonParams {
	return database.RegisterMoveForMonParams{
		Monsterid: monsterID,
		Moveid:    moveID,
	}
}

// tests Querier monsters methods with actual db queries
func (s *DatabaseTestSuite) Test_ListMonsters() {
	ctx := context.Background()
	db := database.New(s.tx)

	monster := NewDatabaseMonsterParams()
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

	monster := NewDatabaseMonsterParams()
	err := db.CreateMonster(ctx, monster)
	s.NoError(err)

	dbMon, err := db.FetchMonster(ctx, monster.ID)
	s.NoError(err)

	s.Equal(monster.ID, dbMon.ID)
}

func (s *DatabaseTestSuite) Test_CreateAndFetchMove() {
	ctx := context.Background()
	db := database.New(s.tx)

	move := NewDatabaseMoveParams()
	err := db.CreateMove(ctx, move)
	s.NoError(err)

	dbMove, err := db.FetchMove(ctx, move.ID)
	s.NoError(err)

	s.Equal(move.ID, dbMove.ID)
}

func (s *DatabaseTestSuite) Test_CreateAndFetchStats() {
	ctx := context.Background()
	db := database.New(s.tx)

	monster := NewDatabaseMonsterParams()
	err := db.CreateMonster(ctx, monster)
	s.NoError(err)

	createParams := NewDatabaseStatParams(monster.ID)
	err = db.CreateStatForMon(ctx, createParams)
	s.NoError(err)

	fetchParams := NewDatabaseFetchStatParams(monster.ID, createParams.Stattype)
	stat, err := db.FetchStat(ctx, fetchParams)
	s.NoError(err)

	s.Equal(createParams.Stattype, stat.Stattype)
}

func (s *DatabaseTestSuite) Test_CreateAndFetchMoveMap() {
	ctx := context.Background()
	db := database.New(s.tx)

	monster := NewDatabaseMonsterParams()
	err := db.CreateMonster(ctx, monster)
	s.NoError(err)

	move := NewDatabaseMoveParams()
	err = db.CreateMove(ctx, move)
	s.NoError(err)

	moveMap := NewRegisterMoveParams(monster.ID, move.ID)
	err = db.RegisterMoveForMon(ctx, moveMap)
	s.NoError(err)

	moves, err := db.FetchMovesForMon(ctx, monster.ID)
	s.NoError(err)
	s.Equal(move.ID, moves[0].Moveid)
}
