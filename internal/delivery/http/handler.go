package http

import (
	"net/http"

	"github.com/aintsashqa/go-observer-service/internal/service"
	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
)

type Handler struct {
	upgrader *websocket.Upgrader
	cookie   sessions.Store
	services *service.Service
}

func NewHandler(cookie sessions.Store, services *service.Service) *Handler {
	return &Handler{
		upgrader: &websocket.Upgrader{},
		cookie:   cookie,
		services: services,
	}
}

func (handler *Handler) Initialize() http.Handler {
	router := chi.NewRouter()

	router.Route("/square", handler.squareRoutes)
	router.Route("/user", handler.userRoutes)

	return router
}
