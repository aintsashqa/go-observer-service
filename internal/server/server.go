package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aintsashqa/go-observer-service/internal/config"
)

type Server struct {
	http *http.Server
}

func NewServer(cfg config.HTTPConfig, handler http.Handler) *Server {
	return &Server{
		http: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler: handler,
		},
	}
}

func (server *Server) Run() error {
	return server.http.ListenAndServe()
}

func (server *Server) Stop(ctx context.Context) error {
	return server.http.Shutdown(ctx)
}
