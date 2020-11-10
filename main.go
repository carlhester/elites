package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var addr = "localhost:8181"

func main() {
	network := flag.Bool("n", false, "network")
	flag.Parse()

	if *network == true {
		log.Printf("starting network on %s", addr)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}
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
