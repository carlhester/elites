package main

import (
	"bufio"
	"os"
	"strconv"
)

func GetInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if input[0:1] == "q" {
		quitGame("\n\nThanks for playing!\n\n")
	}
	// make sure we're getting a valid digit here
	val, _ := strconv.Atoi(input[0:1])
	return val
}
