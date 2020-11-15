package main

import (
	"elite/comms"
	"flag"
	"fmt"
	"os"
)

func main() {
	network := flag.Bool("n", false, "network")
	flag.Parse()

	out := output.NewOutput()
	out.addWriteTo(os.Stdout)

	if *network {
		ip := "0.0.0.0"
		port := 8181
		go comms.Listen(ip, port)
	}
	startGame(out)
}

func startGame(out *output.output) {
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
