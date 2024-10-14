package server

import (
	"fmt"
	"net/http"

	"connectrpc.com/grpchealth"
	"github.com/jchapman63/neo-gkmn/internal/config"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Service interface {
	Register(*http.ServeMux)
	Name() string
}

type Server struct {
	*http.ServeMux
	health *grpchealth.StaticChecker
	config *config.Server
}

func New(cfg *config.Server) (*Server, error) {
	api := &Server{
		ServeMux: http.NewServeMux(),
		health:   grpchealth.NewStaticChecker(),
		config:   cfg,
	}

	api.registerHealthService()

	return api, nil
}

func (s *Server) registerHealthService() {
	s.Handle(grpchealth.NewHandler(s.health))
}

func (s *Server) RegisterService(service Service) {
	service.Register(s.ServeMux)
	s.health.SetStatus(service.Name(), grpchealth.StatusServing)
}

func (s *Server) Serve() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), h2c.NewHandler(s, &http2.Server{}))
}
