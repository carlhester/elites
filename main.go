package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	network := flag.Bool("n", false, "network")
	flag.Parse()

	output := NewOutput()
	output.addWriteTo(os.Stdout)

	if *network {
		go func() {
			addr := &net.TCPAddr{Port: 8181}
			fmt.Printf("listening on %s", addr)
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
				output.addWriteTo(conn)
			}
		}()
	}
	startGame(output)
}

func startGame(output *output) {
	g := &game{
		turn:   1,
		output: output,
	}
	g.Run()
}

func quitGame(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}
