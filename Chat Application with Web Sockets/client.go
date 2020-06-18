package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
type client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn //keeps a ref to the web socket
	// send is a channel on which messages are sent.
	send chan []byte //buffered channel to queue received msgs
	// room is the room this client is chatting in.
	room *room //ref to the room that the client is chatting in.
}

// for init communication on web socket
func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
