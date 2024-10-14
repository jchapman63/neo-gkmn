package gkmn

import (
	"context"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"github.com/jchapman63/neo-gkmn/internal/database"
	"github.com/jchapman63/neo-gkmn/internal/pkg"
	gkmnv1 "github.com/jchapman63/neo-gkmn/internal/service/gkmn/v1"
	"github.com/jchapman63/neo-gkmn/internal/service/gkmn/v1/gkmnv1connect"
)

type GameHandler struct {
	db      database.Querier
	options []connect.HandlerOption
}

// CreateBattle implements gkmnv1connect.GkmnServiceHandler.
func (h *GameHandler) CreateBattle(ctx context.Context, req *connect.Request[gkmnv1.GkmnServiceCreateBattleRequest]) (*connect.Response[gkmnv1.GkmnServiceCreateBattleResponse], error) {
	monsters := req.Msg.GetMonsters()

	var dbMon []*database.Monster
	for _, mon := range monsters {
		mapped := database.Monster{
			ID:     mon.GetId(),
			Name:   mon.GetName(),
			Type:   mon.GetType(),
			Basehp: mon.GetBaseHp(),
		}
		dbMon = append(dbMon, &mapped)
	}

	battle, err := pkg.NewBattle(ctx, h.db, dbMon)
	if err != nil {
		slog.Error("could not create new battle", "err", err)
		return nil, err
	}

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
	h := &GameHandler{db: db}
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
