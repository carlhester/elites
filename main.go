package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
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
			go startGame(conn)
		}
	} else {
		startGame(os.Stdout)
	}

}

func startGame(output io.Writer) {
	g := &game{
		turn:   1,
		output: NewOutput(output),
	}
	g.Run()
}

func quitGame(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}
