package gkmn

import (
	"context"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"github.com/jchapman63/neo-gkmn/internal/battle"
	gkmnv1 "github.com/jchapman63/neo-gkmn/internal/connect/gkmn/v1"
	"github.com/jchapman63/neo-gkmn/internal/connect/gkmn/v1/gkmnv1connect"
	"github.com/jchapman63/neo-gkmn/internal/database"
)

type GameHandler struct {
	db            database.Querier
	options       []connect.HandlerOption
	activeBattles map[string]*battle.Battle
}

type GameServiceOption func(n *GameHandler)

func WithHandlerOptions(opts ...connect.HandlerOption) GameServiceOption {
	return func(g *GameHandler) {
		g.options = opts
	}
}

func NewGameService(db database.Querier, opts ...GameServiceOption) *GameHandler {
	h := &GameHandler{db: db, activeBattles: make(map[string]*battle.Battle)}
	for _, o := range opts {
		o(h)
	}
	return h
}

// implement service interface
func (h *GameHandler) Register(router *http.ServeMux) {
	router.Handle(gkmnv1connect.NewGkmnServiceHandler(h, h.options...))
}

// helps expose to reflection
func (h *GameHandler) Name() string {
	return gkmnv1connect.GkmnServiceName
}

// ListActiveBattles implements gkmnv1connect.GkmnServiceHandler.
func (h *GameHandler) ListActiveBattles(context.Context, *connect.Request[gkmnv1.GkmnServiceActiveBattleListRequest]) (*connect.Response[gkmnv1.GkmnServiceActiveBattleListResponse], error) {
	var battles []string
	for _, b := range h.activeBattles {
		battles = append(battles, b.ID)
	}

	return connect.NewResponse(&gkmnv1.GkmnServiceActiveBattleListResponse{
		BattleIds: battles,
	}), nil
}

// ListBattleMonsters implements gkmnv1connect.GkmnServiceHandler.
// Grabs monsters from an active game
func (h *GameHandler) ListBattleMonsters(ctx context.Context, req *connect.Request[gkmnv1.GkmnServiceListBattleMonsterRequest]) (*connect.Response[gkmnv1.GkmnServiceListBattleMonsterResponse], error) {
	return connect.NewResponse(&gkmnv1.GkmnServiceListBattleMonsterResponse{
		BattleMonsters: h.MapBattleMonsters(req.Msg.GetBattleId()),
		BattleId:       req.Msg.GetBattleId(),
	}), nil
}

// ListMonsters implements gkmnv1connect.GkmnServiceHandler.
func (h *GameHandler) ListMonsters(ctx context.Context, req *connect.Request[gkmnv1.GkmnServiceBaseMonsterListRequest]) (*connect.Response[gkmnv1.GkmnServiceBaseMonsterListResponse], error) {
	mons, err := h.db.ListMonsters(ctx)
	if err != nil {
		return nil, err
	}

	var respMons []*gkmnv1.Monster
	for _, mon := range mons {
		monster := &gkmnv1.Monster{
			Id:     mon.ID,
			Name:   mon.Name,
			Type:   mon.Type,
			BaseHp: mon.Basehp,
		}
		respMons = append(respMons, monster)
	}

	return connect.NewResponse(&gkmnv1.GkmnServiceBaseMonsterListResponse{
		Monsters: respMons,
	}), nil
}

// AttackMonster implements gkmnv1connect.GkmnServiceHandler.
func (h *GameHandler) AttackMonster(ctx context.Context, req *connect.Request[gkmnv1.GkmnServiceAttackMonsterRequest]) (*connect.Response[gkmnv1.GkmnServiceAttackMonsterResponse], error) {
	battleID := req.Msg.GetBattleId()
	battle := h.activeBattles[battleID]

	active := battle.PriorityMon()
	if active != req.Msg.GetActorId() {
		slog.Info("not the mons turn")
		// TODO, should return error saying not monsters turn
		return nil, nil
	}
	actor := battle.Monsters[req.Msg.GetActorId()]
	move := actor.Moves[req.Msg.GetMoveId()]

	battle.Damage(req.Msg.GetVictimId(), *move)

	return connect.NewResponse(&gkmnv1.GkmnServiceAttackMonsterResponse{
		BattleState: &gkmnv1.BattleState{
			BattleMonsters: h.MapBattleMonsters(battleID),
			BattleId:       battleID,
			// TODO , rm hardcoding
			IsOver: false,
		},
	}), nil
}

// CreateBattle implements gkmnv1connect.GkmnServiceHandler.
func (h *GameHandler) CreateBattle(ctx context.Context, req *connect.Request[gkmnv1.GkmnServiceCreateBattleRequest]) (*connect.Response[gkmnv1.GkmnServiceCreateBattleResponse], error) {
	monsterRequests := req.Msg.GetMonIds()

	var monIds []string
	for _, mon := range monsterRequests {
		monIds = append(monIds, mon.Id)
	}

	battle, err := battle.NewBattle(ctx, h.db, monIds)
	if err != nil {
		slog.Error("could not create new battle", "err", err)
		return nil, err
	}

	h.activeBattles[battle.ID] = battle

	respBattleMon := []*gkmnv1.BattleMonster{}
	for _, mon := range battle.Monsters {
		respMon := &gkmnv1.Monster{
			Id:     mon.Monster.ID,
			Type:   mon.Monster.Type,
			Name:   mon.Monster.Name,
			BaseHp: mon.Monster.Basehp,
		}
		moves := []*gkmnv1.Move{}
		for _, move := range mon.Moves {
			m := &gkmnv1.Move{
				Id:    move.ID,
				Name:  move.Name,
				Power: move.Power,
				Type:  move.Type,
			}
			moves = append(moves, m)
		}
		battleMon := &gkmnv1.BattleMonster{
			Monster: respMon,
			LiveHp:  mon.LiveHp,
			Speed:   mon.Speed,
			Moves:   moves,
		}
		respBattleMon = append(respBattleMon, battleMon)
	}

	return connect.NewResponse(&gkmnv1.GkmnServiceCreateBattleResponse{
		Id:             battle.ID,
		BattleMonsters: respBattleMon,
	}), nil
}
