package pkg

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jchapman63/neo-gkmn/internal/database"
)

type Battle struct {
	ID       string
	db       database.Querier
	Monsters []*BattleMon
}

type BattleMon struct {
	Monster *database.Monster
	LiveHp  int32
	Speed   int32
	Moves   []*database.Move
}

func NewBattle(ctx context.Context, db database.Querier, monIDs []string) (*Battle, error) {

	var BattleMonsters []*BattleMon
	for _, id := range monIDs {

		params := database.FetchStatParams{
			Column1: id,
			Column2: "speed",
		}

		speed, err := db.FetchStat(ctx, params)
		if err != nil {
			msg := fmt.Sprintf("failed to fetch %s stats for monID: %s", params.Column2, params.Column1)
			slog.Error(msg, "err", err)
			return nil, err
		}

		mon, err := db.FetchMonster(ctx, id)
		if err != nil {
			msg := fmt.Sprintf("failed to fetch monster by id: %s", id)
			slog.Error(msg, "err", err)
			return nil, err
		}

		battleMon := &BattleMon{
			Monster: &mon,
			LiveHp:  mon.Basehp,
			Speed:   speed.Power,
			Moves:   []*database.Move{},
		}
		movemap, err := db.FetchMovesForMon(ctx, id)
		for _, mapping := range movemap {
			move, err := db.FetchMove(ctx, mapping.Moveid)
			if err != nil {
				return nil, err
			}
			battleMon.Moves = append(battleMon.Moves, &move)

		}
		BattleMonsters = append(BattleMonsters, battleMon)
	}
	return &Battle{
		ID:       uuid.NewString(),
		db:       db,
		Monsters: BattleMonsters,
	}, nil
}

func (b *Battle) Damage(victimID string, move database.Move) {
	for _, mon := range b.Monsters {
		if mon.Monster.ID == victimID {
			mon.LiveHp -= move.Power
		}
	}
}

func (b *Battle) IsOver() bool {
	for _, mon := range b.Monsters {
		if mon.LiveHp <= 0 {
			return true
		}
	}
	return false
}

// TODO : make use of priority queue instead
// pqueue will change the structure of Battle
func (b *Battle) TurnDecider() *BattleMon {
	maxMon := b.Monsters[0]
	for _, mon := range b.Monsters {
		if mon.Speed > maxMon.Speed {
			maxMon = mon
		}
	}
	return maxMon
}
