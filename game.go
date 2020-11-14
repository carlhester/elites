package main

import (
	"fmt"
	"os"
	"time"
)

type game struct {
	players []*player
	turn    int
	output  *output
}

func (g game) Run() {
	/*
		for x := 0; x < 100; x++ {
			fmt.Printf("clients: %+v", g.output.writeTo)
			g.output.Add(fmt.Sprintf("%d", x))
			g.output.Render()
			time.Sleep(1 * time.Second)
		}
	*/

	g.output.Clear()

	chars := LoadElites()
	p1 := &player{Elite: CharacterSelectMenu(1, chars, g.output), output: g.output}
	p2 := &player{Elite: CharacterSelectMenu(2, chars, g.output), output: g.output}
	p1.enemy = p2
	p2.enemy = p1

	g.players = []*player{p1, p2}
	for {
		for _, p := range g.players {
			g.output.Clear()
			g.output.Add("")
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
	g.output.Add("================")
	for _, p := range g.players {
		g.output.Add(fmt.Sprintf("%s:%d", p.Name, p.Hp))
	}
	g.output.Add("================")
	g.output.Render()
}

func (g game) CheckEnd() {
	for i := range g.players {
		if g.players[i].Hp <= 0 {
			g.output.Add(fmt.Sprintf("%s has been defeated! Good Game!", g.players[i].Name))
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
