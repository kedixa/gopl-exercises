package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// convert to tcp Conn
	tcp := conn.(*net.TCPConn)
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, tcp) // NOTE: ignoring errors
		tcp.CloseRead()
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(tcp, os.Stdin)
	tcp.CloseWrite() // close
	<-done           // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
