package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

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
	fmt.Println()
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

func Clear() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

func CharacterSelectMenu(player int, chars Characters) Elite {
	for i, c := range chars.Elites {
		fmt.Println(i, c)
	}
	choice := GetInput()
	return chars.Elites[choice]
}

func MainMenu() {
	Clear()

	for i := 0; i < 13; i++ {
		fmt.Printf("*")
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("\n*   ELITE   *\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("*           *\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("*  MONSTER  *\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("*           *\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("*   CARDS   *\n")
	for i := 0; i < 13; i++ {
		fmt.Printf("*")
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("\n\nPress enter to continue...")
	_ = GetInput()

}

func (g game) Run() {
	MainMenu()
	chars := LoadElites()
	player1 := CharacterSelectMenu(1, chars)
	player2 := CharacterSelectMenu(2, chars)
	p1 := &player{Elite: player1}
	p2 := &player{Elite: player2}
	p1.enemy = p2
	p2.enemy = p1

	g.players = []*player{p1, p2}
	//g.players = append(players, p1, p2)
	for turn := 0; turn < 10; turn++ {
		for _, p := range g.players {
			g.showStatus()
			g.players[g.turn-1].showOpts()

			validInput := false
			for validInput != true {
				err := p.handleInput(GetInput())
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
