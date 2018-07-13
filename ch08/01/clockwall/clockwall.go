package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type placeHost struct {
	Place  string
	Host   string
	Conn   net.Conn
	Reader *bufio.Reader
}

func main() {
	var ph []placeHost
	// defer conn.Close
	defer func() {
		for _, p := range ph {
			p.Conn.Close()
		}
	}()

	// deal with args
	for _, arg := range os.Args[1:] {
		eq := strings.IndexByte(arg, '=')
		if eq == -1 {
			// illegal
			continue
		}
		// split arg and connect to server
		place := arg[:eq]
		host := arg[eq+1:]
		conn, err := net.Dial("tcp", host)
		if err != nil {
			log.Fatal(err)
			continue
		}
		ph = append(ph, placeHost{place, host, conn, bufio.NewReader(conn)})
	}

	for {
		// show the time of all places
		for _, p := range ph {
			line, err := p.Reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
				continue
			}
			fmt.Printf("The time of %s is: %s", p.Place, string(line))
		}
	}
}
