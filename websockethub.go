package gop5js

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var wshub = newWebsocketHub()

type hub struct {
	connections map[*websocket.Conn]*websocket.Conn
}

func newWebsocketHub() *hub {
	return &hub{
		connections: make(map[*websocket.Conn]*websocket.Conn),
	}
}

func (h *hub) addConnection(c *websocket.Conn) {
	h.connections[c] = c
}

func (h *hub) removeConnection(c *websocket.Conn) {
	delete(h.connections, c)
}

func (h *hub) writeJSON(v interface{}) error {
	for c := range h.connections {
		err := c.WriteJSON(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *hub) writeMessage(messageType int, data []byte) error {
	for c := range h.connections {
		err := c.WriteMessage(messageType, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func wsHandleFunc(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	defer wshub.removeConnection(conn)
	wshub.addConnection(conn)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("error when reading from browser:", err)
			fmt.Println("Program exit.")
			os.Exit(1)
		}
		//fmt.Println(string(message))
		nextFrame(message)
	}
}
