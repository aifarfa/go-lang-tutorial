package main

import (
	"fmt"
	"net"
)

type Chat struct {
	connections map[net.Conn]string
	busy        chan bool
}

func NewChat() *Chat {
	return &Chat{
		connections: make(map[net.Conn]string, 10),
		busy:        make(chan bool, 1),
	}
}

func (list *Chat) Add(conn net.Conn, name string) {
	list.busy <- true
	// list.connections = append(list.connections, conn) slice
	list.connections[conn] = name
	<-list.busy
}

func (list *Chat) Remove(conn net.Conn) {
	list.busy <- true
	delete(list.connections, conn)
	<-list.busy
}

func (list *Chat) Boardcast(msg string) {
	list.busy <- true
	for conn, _ := range list.connections {
		fmt.Fprintf(conn, "%s\n", msg)
	}
	<-list.busy
}

func (list *Chat) Send(user, msg string) {
	list.busy <- true
	for conn, name := range list.connections {
		if user != name {
			fmt.Fprintf(conn, "%s say: %s\n", user, msg)
		}
	}
	<-list.busy
}
