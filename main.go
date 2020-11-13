package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var clients []io.Writer
	clients = append(clients, os.Stdout)

	network := flag.Bool("n", false, "network")
	flag.Parse()

	if *network {
		addr := &net.TCPAddr{Port: 8181}
		log.Printf("starting network on %s", addr)
		listener, err := net.ListenTCP("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}
			fmt.Printf("Connect: %s...\n", conn.RemoteAddr())
			time.Sleep(3 * time.Second)
			clients = append(clients, conn)
			for _, client := range clients {
				fmt.Fprintf(client, "hello!\n")
			}
			go startGame(clients)
		}
	} else {
		startGame(clients)
	}

}

func startGame(clients []io.Writer) {
	g := &game{
		turn:    1,
		outputs: NewOutputs(clients),
	}
	g.Run()
}

func quitGame(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}
