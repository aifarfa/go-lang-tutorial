package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var chat *Chat = NewChat()

func main() {
	ln, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Println("start listening at :2222")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	regist(conn, scanner)
	echo(conn, scanner)
}

func regist(conn net.Conn, scanner *bufio.Scanner) {
	fmt.Fprint(conn, "What is your name? ")
	scanner.Scan()
	name := scanner.Text()
	chat.Add(conn, name)
	chat.Boardcast("new user joined: " + chat.connections[conn])
	log.Println(name, "joined the chat. total users:", len(chat.connections))
}

func echo(conn net.Conn, scanner *bufio.Scanner) {
  user := chat.connections[conn]
	for scanner.Scan() {
		msg := scanner.Text()
		chat.Send(user, msg)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// when disconnected
	chat.Boardcast(user + " has left the chat")
	chat.Remove(conn)
	log.Println(user, "has left. total users:", len(chat.connections))
}

func Timer(conn net.Conn) {
	for counter := 30; counter >= 0; counter-- {
		time.Sleep(1 * time.Second)
		if counter < 5 || counter%10 == 0 {
			fmt.Fprintf(conn, "disconnect in: %d seconds \n", counter)
		}
	}
	conn.Close()
}
