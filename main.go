package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Bool("n", false, "network")
	flag.Parse()

	g := &game{
		turn:   1,
		output: NewOutput(os.Stdout),
	}
	g.Run()
}

func quitGame(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}
