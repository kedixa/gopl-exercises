package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	Name string
	C    chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				// skip if cli.C is not ready
				select {
				case cli.C <- msg:
				default:
				}
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.C)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string, 100) // outgoing client messages with buffer
	go clientWriter(conn, ch)

	// let the client choose a name
	input := bufio.NewScanner(conn)
	ch <- "Please input your name: "
	input.Scan()
	who := input.Text()

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{who, ch}

	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- client{who, ch}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
