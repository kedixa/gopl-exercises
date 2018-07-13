package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	var text = make(chan string)
	// read inputs and send to chan text
	go func() {
		for input.Scan() {
			text <- input.Text()
		}
		log.Print("Closed")
	}()
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case t := <-text:
			go echo(c, t, 1*time.Second)
			ticker.Stop()
			ticker = time.NewTicker(10 * time.Second)
		case <-ticker.C:
			// timeout
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
