package app

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/aintsashqa/go-observer-service/internal/config"
	"github.com/aintsashqa/go-observer-service/internal/delivery/http"
	"github.com/aintsashqa/go-observer-service/internal/server"
	"github.com/aintsashqa/go-observer-service/internal/service"
	"github.com/gorilla/sessions"
)

func Run() {
	cfg, err := config.InitConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	services := service.NewService()
	cookie := sessions.NewCookieStore([]byte(cfg.App.CookieKey))
	handler := http.NewHandler(cookie, services)

	server := server.NewServer(cfg.HTTP, handler.Initialize())
	go func() {
		if err := server.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Starting server on %s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, os.Kill)
	<-signalCh
	if err := server.Stop(ctx); err != nil {
		log.Fatal(err)
	}
	log.Printf("Graceful shutdown successfully done")
}
