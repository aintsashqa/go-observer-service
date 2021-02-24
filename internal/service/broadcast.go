package service

import (
	"log"

	"github.com/gorilla/websocket"
)

type BroadcastServiceImpl struct {
	channel     chan []byte
	connections map[*websocket.Conn]bool
}

func NewBroadcastServiceImpl() *BroadcastServiceImpl {
	return &BroadcastServiceImpl{
		channel:     make(chan []byte),
		connections: make(map[*websocket.Conn]bool),
	}
}

func (broadcast *BroadcastServiceImpl) write(connection *websocket.Conn) {
	defer broadcast.removeClient(connection)

	for {
		select {
		case message, ok := <-broadcast.channel:
			if !ok {
				connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			writer, err := connection.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Print(err)
				return
			}

			writer.Write(message)

			if err := writer.Close(); err != nil {
				log.Print(err)
				return
			}
		}
	}
}

func (broadcast *BroadcastServiceImpl) read(connection *websocket.Conn) {
	defer broadcast.removeClient(connection)

	for {
		_, message, err := connection.ReadMessage()
		if err != nil {
			log.Print(err)
			break
		}

		broadcast.channel <- message
	}
}

func (broadcast *BroadcastServiceImpl) removeClient(connection *websocket.Conn) {
	connection.Close()
	delete(broadcast.connections, connection)
}

func (broadcast *BroadcastServiceImpl) AddClient(connection *websocket.Conn) {
	broadcast.connections[connection] = true

	go broadcast.write(connection)
	go broadcast.read(connection)
}
