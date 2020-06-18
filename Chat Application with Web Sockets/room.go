package main

type room struct {
	forward chan []byte
}

// forward is a channel that holds incoming messages
// that should be forwarded to the other clients.
