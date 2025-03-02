package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"github.com/jchapman63/neo-gkmn/internal/config"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// connect services will implement this
type Service interface {
	Register(*http.ServeMux)
	Name() string
}

type Server struct {
	*http.ServeMux
	reflection *Reflector
	health     *grpchealth.StaticChecker
	config     *config.Server
}

func New(cfg *config.Server) (*Server, error) {
	api := &Server{
		ServeMux:   http.NewServeMux(),
		reflection: NewReflector(),
		health:     grpchealth.NewStaticChecker(),
		config:     cfg,
	}

	api.registerHealthService()
	api.registerReflectionService()

	return api, nil
}

func (s *Server) registerReflectionService() {
	s.Handle(grpcreflect.NewHandlerV1Alpha(grpcreflect.NewReflector(s.reflection)))
}

func (s *Server) registerHealthService() {
	s.Handle(grpchealth.NewHandler(s.health))
}

func (s *Server) RegisterService(service Service) {
	service.Register(s.ServeMux)
	s.reflection.AddService(service.Name())
	s.health.SetStatus(service.Name(), grpchealth.StatusServing)
}

func (s *Server) Serve() error {
	slog.Info("starting server", "port", s.config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), h2c.NewHandler(s, &http2.Server{}))
}
