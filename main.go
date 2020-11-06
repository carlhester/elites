package main

import (
	"fmt"
	"os"
)

func main() {
	g := &game{
		turn:   1,
		output: &Output{},
	}
	g.Run()
}

func quitGame(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}
