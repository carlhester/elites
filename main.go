package main

import (
	"fmt"
	"os"
)

func quitGame(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}

func main() {
	g := &game{turn: 1}
	g.Run()
}

func Clear() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}
