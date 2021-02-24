package service

import (
	"github.com/gorilla/websocket"
)

type (
	BroadcastServiceInterface interface {
		AddClient(*websocket.Conn)
	}

	Service struct {
		Broadcast BroadcastServiceInterface
	}
)

func NewService() *Service {
	return &Service{
		Broadcast: NewBroadcastServiceImpl(),
	}
}
