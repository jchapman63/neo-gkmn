package battle

import (
	"container/heap"
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jchapman63/neo-gkmn/internal/database"
)

type Battle struct {
	ID       string
	db       database.Querier
	Monsters map[string]*BattleMon
	Queue    *PriorityQueue
}

type BattleMon struct {
	Monster *database.Monster
	LiveHp  int32
	Speed   int32
	Moves   map[string]*database.Move
}

func NewBattle(ctx context.Context, db database.Querier, monIDs []string) (*Battle, error) {

	battleMonsters := map[string]*BattleMon{}
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
			Moves:   map[string]*database.Move{},
		}
		movemap, err := db.FetchMovesForMon(ctx, id)
		if err != nil {
			slog.Error("failed to fetch moves for mon", "err", err)
			return nil, err
		}
		for _, mapping := range movemap {
			move, err := db.FetchMove(ctx, mapping.Moveid)
			if err != nil {
				return nil, err
			}
			battleMon.Moves[move.ID] = &move
		}
		battleMonsters[id] = battleMon
	}

	return &Battle{
		ID:       uuid.NewString(),
		db:       db,
		Monsters: battleMonsters,
		Queue:    initQueue(battleMonsters),
	}, nil
}

func initQueue(mons map[string]*BattleMon) *PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for _, m := range mons {
		pq.Push(m)
	}
	return &pq
}

func (b *Battle) Damage(victimID string, move database.Move) {
	b.Monsters[victimID].LiveHp -= move.Power
}

func (b *Battle) IsOver() bool {
	for _, mon := range b.Monsters {
		if mon.LiveHp <= 0 {
			return true
		}
	}
	return false
}

func (b *Battle) PriorityMon() string {
	if b.Queue.Len() == 0 {
		b.Queue = initQueue(b.Monsters)
	}
	item := b.Queue.Pop()

	return string(item.(*Item).monID)
}
