package gkmn

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	gkmnv1 "github.com/jchapman63/neo-gkmn/internal/connect/gkmn/v1"
	"github.com/jchapman63/neo-gkmn/internal/connect/gkmn/v1/gkmnv1connect"
	"github.com/jchapman63/neo-gkmn/internal/database"
	"github.com/jchapman63/neo-gkmn/internal/pkg"
)

var ATTACK string = "ATTACK"

type ActionRequest struct {
	battleID  string
	action    string
	request   *connect.Request[gkmnv1.GkmnServiceAttackMonsterRequest]
	completed bool
}

type AlterRequest func()

func NewActionRequest(battleID string, action string, req *connect.Request[gkmnv1.GkmnServiceAttackMonsterRequest], opts ...func(*ActionRequest)) *ActionRequest {
	a := &ActionRequest{
		battleID:  battleID,
		action:    action,
		request:   req,
		completed: false,
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
				return
			}
			battle.Damage(actReq.request.Msg.GetVictimId(), move)
		}
		actReq.completed = true
	}
}
