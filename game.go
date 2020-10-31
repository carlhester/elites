package main

import (
	"fmt"
	"os"
)

type game struct {
	players []*player
	turn    int
}

func (g game) Run() {
	MainMenu()
	Clear()

	chars := LoadElites()
	p1 := &player{Elite: CharacterSelectMenu(1, chars)}
	p2 := &player{Elite: CharacterSelectMenu(2, chars)}
	p1.enemy = p2
	p2.enemy = p1

	g.players = []*player{p1, p2}
	for turn := 0; turn < 10; turn++ {
		for _, p := range g.players {
			Clear()
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
			fmt.Printf("%s has been defeated! Good Game!", g.players[i].Name)
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
