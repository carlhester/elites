package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	ip := &net.TCPAddr{IP: net.ParseIP("127.0.0.1")}
	port := &net.TCPAddr{Port: 8181}
	conn, err := net.DialTCP("tcp", ip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	conn.SetKeepAlive(true)
	conn.SetLinger(10)
	conn.SetNoDelay(false)

	//reader := bufio.NewScanner(conn)
	for {
		done := make(chan struct{})
		go func() {
			io.Copy(os.Stdout, conn) // NOTE: ignoring errors
			done <- struct{}{}       // signal the main goroutine
		}()
		time.Sleep(1 * time.Second)
	}
}
