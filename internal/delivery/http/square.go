package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

func (handler *Handler) squareRoutes(router chi.Router) {
	router.Get("/", handler.GetSquare)
	router.Get("/observe", handler.ObserveSquare)
}

func (handler *Handler) socket(response http.ResponseWriter, request *http.Request) (*websocket.Conn, error) {
	handler.upgrader.CheckOrigin = func(_ *http.Request) bool { return true }
	return handler.upgrader.Upgrade(response, request, nil)
}

func (handler *Handler) GetSquare(response http.ResponseWriter, request *http.Request) {
	session, _ := handler.cookie.Get(request, "session-name")
	isCurrentUserAdmin := false
	if value, ok := session.Values["is_current_user_admin"]; ok {
		isCurrentUserAdmin = value.(bool)
	}

	handler.ResponseHTML(response, viewGetSquareTemplateFilename, struct {
		IsAdmin bool
	}{
		IsAdmin: isCurrentUserAdmin,
	})
}

func (handler *Handler) ObserveSquare(response http.ResponseWriter, request *http.Request) {
	socket, err := handler.socket(response, request)
	if err != nil {
		log.Print(err)
		return
	}

	handler.services.Broadcast.AddClient(socket)
}
