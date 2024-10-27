package gkmn

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"github.com/jchapman63/neo-gkmn/internal/database"
	"github.com/jchapman63/neo-gkmn/internal/pkg"
	gkmnv1 "github.com/jchapman63/neo-gkmn/internal/service/gkmn/v1"
	"github.com/jchapman63/neo-gkmn/internal/service/gkmn/v1/gkmnv1connect"
)

var ATTACK string = "ATTACK"

type ActionRequest struct {
	battleID string
	action   string
	request  *connect.Request[gkmnv1.GkmnServiceAttackMonsterRequest]
}

type AlterRequest func()

func NewActionRequest(battleID string, action string, req *connect.Request[gkmnv1.GkmnServiceAttackMonsterRequest], opts ...func(*ActionRequest)) *ActionRequest {
	a := &ActionRequest{
		battleID: battleID,
		action:   action,
		request:  req,
	}

	for _, o := range opts {
		o(a)
	}

	return a
}

type GameHandler struct {
	db            database.Querier
	options       []connect.HandlerOption
	activeBattles map[string]*pkg.Battle
	c             chan *ActionRequest
}

// ListBattleMonsters implements gkmnv1connect.GkmnServiceHandler.
func (h *GameHandler) ListBattleMonsters(context.Context, *connect.Request[gkmnv1.GkmnServiceListBattleMonsterRequest]) (*connect.Response[gkmnv1.GkmnServiceListBattleMonsterResponse], error) {
	panic("unimplemented")
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
	battleId := req.Msg.GetBattleId()

	actReq := NewActionRequest(battleId, ATTACK, req)
	h.c <- actReq

	return connect.NewResponse(&gkmnv1.GkmnServiceAttackMonsterResponse{}), nil
}

// CreateBattle implements gkmnv1connect.GkmnServiceHandler.
func (h *GameHandler) CreateBattle(ctx context.Context, req *connect.Request[gkmnv1.GkmnServiceCreateBattleRequest]) (*connect.Response[gkmnv1.GkmnServiceCreateBattleResponse], error) {
	monsterRequests := req.Msg.GetMonIds()

	var monIds []string
	for _, mon := range monsterRequests {
		monIds = append(monIds, mon.Id)
	}

	battle, err := pkg.NewBattle(ctx, h.db, monIds)
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

type GameServiceOption func(n *GameHandler)

func WithHandlerOptions(opts ...connect.HandlerOption) GameServiceOption {
	return func(g *GameHandler) {
		g.options = opts
	}
}

func NewGameService(db database.Querier, opts ...GameServiceOption) *GameHandler {
	h := &GameHandler{db: db, c: make(chan *ActionRequest), activeBattles: make(map[string]*pkg.Battle)}
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

// Listens for actions to be taken in a battle
func (h *GameHandler) Listen(ctx *context.Context) {
	for {
		actReq := <-h.c
		action := actReq.action
		switch action {
		case ATTACK:
			battle := h.activeBattles[actReq.request.Msg.GetBattleId()]
			move, err := h.db.FetchMove(*ctx, actReq.request.Msg.GetMoveId())
			if err != nil {
				msg := fmt.Sprintf("unable to fetch moveId: %s", actReq.request.Msg.GetMoveId())
				slog.Error(msg, "err", err)
			}
			battle.Damage(actReq.request.Msg.GetVictimId(), move)
			fmt.Println(battle.Monsters[0].LiveHp) // TODO remove
		}
	}
}
