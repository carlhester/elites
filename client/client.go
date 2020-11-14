package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ip := &net.TCPAddr{IP: net.ParseIP("127.0.0.1")}
	port := &net.TCPAddr{Port: 8181}
	conn, err := net.DialTCP("tcp", ip, port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("-> " + message)
	}
}
