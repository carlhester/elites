package main

import (
	"fmt"
	"os"
	"time"
)

type game struct {
	players []*player
	turn    int
	output  *Output
}

func (g game) Run() {
	g.output.Clear()

	chars := LoadElites()
	p1 := &player{Elite: CharacterSelectMenu(1, chars)}
	p2 := &player{Elite: CharacterSelectMenu(2, chars)}
	p1.enemy = p2
	p2.enemy = p1

	g.players = []*player{p1, p2}
	for {
		for _, p := range g.players {
			g.output.Clear()
			g.showStatus()
			validInput := false
			for validInput != true {
				g.players[g.turn-1].showMoves()
				g.players[g.turn-1].MovePrompt()

				err := p.handleInput(GetInput())
				if err != nil {
					fmt.Println(err.Error())
					time.Sleep(2 * time.Second)
					continue
				}
				validInput = true
			}

			p.doNextMove()
			g.CheckEnd()
			g.changeTurns()
		}
	}
}

func (g game) showStatus() {
	fmt.Println("================")
	for _, g := range g.players {
		fmt.Println(g.Name, ":", g.Hp)
	}
	fmt.Println()
}

func (g game) CheckEnd() {
	for i := range g.players {
		if g.players[i].Hp <= 0 {
			fmt.Printf("\n\n%s has been defeated! Good Game!\n\n", g.players[i].Name)
			os.Exit(0)
		}

	}
}

func (g *game) changeTurns() {
	if g.turn == 1 {
		g.turn = 2
		return
	}
	g.turn = 1
}
