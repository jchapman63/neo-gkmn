package gkmn

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/jchapman63/neo-gkmn/internal/database"
	gkmnv1 "github.com/jchapman63/neo-gkmn/internal/service/gkmn/v1"
	"github.com/jchapman63/neo-gkmn/internal/service/gkmn/v1/gkmnv1connect"
)

type GameHandler struct {
	db      database.Querier
	options []connect.HandlerOption
}

// CreateBattle implements gkmnv1connect.GkmnServiceHandler.
func (h *GameHandler) CreateBattle(context.Context, *connect.Request[gkmnv1.NewBattleRequest]) (*connect.Response[gkmnv1.NewBattleResponse], error) {
	panic("unimplemented")
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
