package pkg

import (
	"github.com/google/uuid"
	"github.com/jchapman63/neo-gkmn/internal/database"
)

type Battle struct {
	Monsters []*BattleMon
}

type BattleMon struct {
	Monster *database.Monster
	LiveHp int32
}

func (b *Battle) Damage(victim uuid.UUID, move database.Move) {
	for _, mon := range b.Monsters {
		if mon.Monster.ID.String() == victim.String() {
			mon.LiveHp -= move.Power

		}
	}
}

