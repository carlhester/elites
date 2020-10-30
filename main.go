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

type game struct {
	players []*player
	turn    int
}

func (g game) showStatus() {
	fmt.Println("================")
	for _, g := range g.players {
		fmt.Println(g.Name, ":", g.Hp)
	}
	fmt.Println()
}

func (g *game) changeTurns() {
	if g.turn == 1 {
		g.turn = 2
		return
	}
	g.turn = 1
}

func (g game) GetInput() int {
	fmt.Printf("%s's command: ", g.players[g.turn-1].Name)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if input[0:1] == "q" {
		quitGame("Thanks for playing!")
	}
	val, _ := strconv.Atoi(input[0:1])
	return val
}

type player struct {
	Elite
	nextMove move
	enemy    *player
}

func (p *player) doNextMove() {
	if p.nextMove.MoveType == "heal" {
		p.Hp += p.nextMove.Value
		return
	}
	if p.nextMove.MoveType == "attack" {
		p.enemy.Hp -= p.nextMove.Value
		return
	}
}

func (p player) showOpts() {
	for i := range p.Moves {
		if p.Moves[i].Uses == -1 {
			fmt.Printf("[%d]  %s\n", i, p.Moves[i].Name)
		} else {
			fmt.Printf("[%d]  %s (Uses: %d)\n", i, p.Moves[i].Name, p.Moves[i].Uses)
		}

	}
}

func (p *player) handleInput(cmd int) error {
	if cmd > len(p.Moves) {
		return fmt.Errorf("Invalid Selection")
	}
	if p.Moves[cmd].Uses == 0 {
		return fmt.Errorf("No Uses left")
	}
	if p.Moves[cmd].Uses > 0 {
		p.Moves[cmd].Uses -= 1
	}
	p.nextMove = p.Moves[cmd]
	return nil
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

func (g game) Run() {

	chars := LoadElites()

	var players []*player
	p1 := &player{Elite: chars.Elites[0]}
	p2 := &player{Elite: chars.Elites[1]}
	p1.enemy = p2
	p2.enemy = p1

	g.players = append(players, p1, p2)
	for turn := 0; turn < 10; turn++ {
		for _, p := range g.players {
			g.showStatus()
			g.players[g.turn-1].showOpts()

			validInput := false
			for validInput != true {
				err := p.handleInput(g.GetInput())
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				validInput = true
			}

			p.doNextMove()
			g.changeTurns()
		}
	}
}
