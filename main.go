package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	out := os.Stdout
	network := flag.Bool("n", false, "network")
	flag.Parse()

	if *network == true {
		fmt.Println("starting network")
		//startlistener
		//getconn
		//out := conn
	}

	g := &game{
		turn:   1,
		output: NewOutput(out),
	}
	g.Run()
}

func quitGame(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}
