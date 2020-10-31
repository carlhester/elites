package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Elite struct {
	Name  string `yaml:"Name"`
	Hp    int    `yaml:"Hp"`
	Moves []move `yaml:"Moves"`
}

type Characters struct {
	Elites []Elite `yaml:"Elites"`
}

func LoadElites() Characters {
	var chars Characters
	source, err := ioutil.ReadFile("./elites.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &chars)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return chars
}

func GetInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if input[0:1] == "q" {
		quitGame("Thanks for playing!")
	}
	// make sure we're getting a valid digit here
	val, _ := strconv.Atoi(input[0:1])
	return val
}

type move struct {
	Name     string `yaml:"Name"`
	Value    int    `yaml:"Value"`
	MoveType string `yaml:"MoveType"`
	Uses     int    `yaml:"Uses"`
}

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
