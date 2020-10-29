package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type move struct {
	name  string
	value int
    moveType string
}

type player struct {
	name     string
	hp       int
	moves    []move
	nextMove move
    enemy *player
}

func (p *player) doNextMove() {
    if p.nextMove.moveType == "heal" {
       p.hp += p.nextMove.value
       return
    }
    if p.nextMove.moveType == "attack" {
        p.enemy.hp -= p.nextMove.value
        return
    }
}

type game struct {
	players []*player
	turn    int
}

func (g game) showStatus() {
    fmt.Println("================")
	for _, g := range g.players {
        fmt.Println(g.name,":", g.hp)
	}
}

func (p player) showOpts() {
	for i := range p.moves {
		fmt.Println(i, p.moves[i].name)
	}
}

func (p *player) handleInput(cmd int) {
    p.nextMove = p.moves[cmd]
}

func main() {
	var players []*player
    healme := move{name: "healme", value: 3, moveType: "heal"}
    attack := move{name: "attack", value: 2, moveType: "attack"}
    p1 := &player{name: "P1", hp: 100, moves: []move{healme, attack}}
    p2 := &player{name: "P2", hp: 100, moves: []move{healme, attack}, enemy: p1}
    p1.enemy = p2

	g := &game{}
	g.players = append(players, p1, p2)

	for turn := 0; turn < 10; turn++ {
		for _, p := range g.players {
			g.showStatus()
			p.showOpts()
			p.handleInput(p.GetInput())
            p.doNextMove()
		}
	}
}

func (p player)GetInput() int {
    fmt.Println(p.name, "'s turn:  ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	inputChar := input[0:1]
	val, _ := strconv.Atoi(inputChar)
	return val
}
