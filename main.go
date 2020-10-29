package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type game struct {
	players []*player
	turn    int
}

func (g game) showStatus() {
	fmt.Println("================")
	for _, g := range g.players {
		fmt.Println(g.name, ":", g.hp)
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
    fmt.Printf("%s's command: ", g.players[g.turn-1].name)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
    if input[0:1] == "q" { 
        quitGame("Thanks for playing!")
    } 
    val, _ := strconv.Atoi(input[0:1])
	return val
}

type player struct {
	name     string
	hp       int
	moves    []move
	nextMove move
	enemy    *player
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

func (p player) showOpts() {
	for i := range p.moves {
        if p.moves[i].uses == -1 { 
            fmt.Printf("[%d]  %s\n", i, p.moves[i].name)
        } else {
            fmt.Printf("[%d]  %s (uses: %d)\n", i, p.moves[i].name, p.moves[i].uses)
        }

	}
}

func (p *player) handleInput(cmd int) error {
    if  cmd > len(p.moves) { 
        return fmt.Errorf("Invalid Selection")
    }
    if p.moves[cmd].uses == 0 { 
        return fmt.Errorf("No uses left")
    }
    if p.moves[cmd].uses > 0 { 
        p.moves[cmd].uses -= 1 
    }
	p.nextMove = p.moves[cmd]
    return nil
}

type move struct {
	name     string
	value    int
	moveType string
    uses int
}

func quitGame(msg string) { 
    fmt.Println(msg)
    os.Exit(0)
}


func main() {
	var players []*player
    heal1 := move{name: "heal1", value: 1, moveType: "heal", uses: -1}
    heal2 := move{name: "heal2", value: 2, moveType: "heal", uses: -1}
    heal3 := move{name: "heal3", value: 3, moveType: "heal", uses: -1}
    heal4 := move{name: "heal4", value: 4, moveType: "heal", uses: 2}
    attack := move{name: "attack", value: 2, moveType: "attack", uses: -1}
	p1 := &player{name: "P1", hp: 100, moves: []move{heal1, heal2, heal3, heal4, attack}}
	p2 := &player{name: "P2", hp: 100, moves: []move{heal1, attack}, enemy: p1}
	p1.enemy = p2

    var validInput bool
    g := &game{turn: 1}
	g.players = append(players, p1, p2)
	for turn := 0; turn < 10; turn++ {
		for _, p := range g.players {
			g.showStatus()
            g.players[g.turn-1].showOpts()

            validInput = false
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
