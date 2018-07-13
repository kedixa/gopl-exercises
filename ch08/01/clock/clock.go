package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port int

func main() {
	// parse arg port
	flag.IntVar(&port, "port", 8000, "port")
	flag.Parse()
	listenner, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
