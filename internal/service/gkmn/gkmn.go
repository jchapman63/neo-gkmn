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
	request  *connect.Request[gkmnv1.GkmnServiceBattleAttackMonsterRequest]
}

type AlterRequest func()

func NewActionRequest(battleID string, action string, req *connect.Request[gkmnv1.GkmnServiceBattleAttackMonsterRequest], opts ...func(*ActionRequest)) *ActionRequest {
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

// AttackMonster implements gkmnv1connect.GkmnServiceHandler.
func (h *GameHandler) AttackMonster(ctx context.Context, req *connect.Request[gkmnv1.GkmnServiceBattleAttackMonsterRequest]) (*connect.Response[gkmnv1.GkmnServiceCreateBattleResponse], error) {
	battleId := req.Msg.GetBattleId()

	actReq := NewActionRequest(battleId, ATTACK, req)
	h.c <- actReq

	return connect.NewResponse(&gkmnv1.GkmnServiceCreateBattleResponse{}), nil
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

	return connect.NewResponse(&gkmnv1.GkmnServiceCreateBattleResponse{
		Id: battle.ID,
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

// listens for new games to be created
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
