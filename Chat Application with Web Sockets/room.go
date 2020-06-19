package main

type room struct {
	forward chan []byte // forward is a channel that holds incoming messages that should be forwarded to the other clients.

	join    chan *client     //for clients that are joining
	leave   chan *client     //for leaving clients
	clients map[*client]bool //list of curr clients
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			for c := range r.clients {
				c.send <- msg
			}
		}
	}
}
