package pkg

import (
	"context"

	"github.com/google/uuid"
	"github.com/jchapman63/neo-gkmn/internal/database"
)

type Battle struct {
	db       database.Querier
	Monsters []*BattleMon
}

type BattleMon struct {
	Monster *database.Monster
	LiveHp  int32
	Speed   int32
}

func NewBattle(ctx context.Context, db database.Querier, monsters []database.Monster) (*Battle, error) {
	var BattleMonsters []*BattleMon
	for _, mon := range monsters {
		params := database.FetchStatParams{
			Column1: mon.ID,
			Column2: "Speed",
		}
		speed, err := db.FetchStat(ctx, params)
		if err != nil {
			return nil, err
		}
		battleMon := &BattleMon{
			Monster: &mon,
			LiveHp:  mon.Basehp,
			Speed:   speed.Power,
		}
		BattleMonsters = append(BattleMonsters, battleMon)
	}
	return &Battle{
		db:       db,
		Monsters: BattleMonsters,
	}, nil
}

func (b *Battle) Damage(victim uuid.UUID, move database.Move) {
	for _, mon := range b.Monsters {
		if mon.Monster.ID.String() == victim.String() {
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

func (b *Battle) TurnDecider() *BattleMon {
	maxMon := b.Monsters[0]
	for _, mon := range b.Monsters {
		if mon.Speed > maxMon.Speed {
			maxMon = mon
		}
	}
	return maxMon
}
