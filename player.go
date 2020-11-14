package main

import (
	"fmt"
	"time"
)

type player struct {
	Elite
	nextMove move
	enemy    *player
	Stunned  int
	output   *output
}

func (p *player) doNextMove() {
	if p.Stunned > 0 {
		p.output.Add(fmt.Sprintf("%s tries to use %s but they are STUNNED and cannot move!!", p.Name, p.nextMove.Name))
		p.output.Render()
		p.Stunned -= 1
		time.Sleep(2 * time.Second)
		return
	}

	if p.nextMove.MoveType == "heal" {
		p.Hp += p.nextMove.Value
		p.output.Add(fmt.Sprintf("%s uses %s and heals for %d", p.Name, p.nextMove.Name, p.nextMove.Value))
		p.output.Render()
		time.Sleep(2 * time.Second)
		return
	}
	if p.nextMove.MoveType == "attack" {
		p.enemy.Hp -= p.nextMove.Value
		p.output.Add(fmt.Sprintf("%s uses %s! %s takes %d damage.", p.Name, p.nextMove.Name, p.enemy.Name, p.nextMove.Value))
		p.output.Render()
		time.Sleep(2 * time.Second)
		return
	}
	if p.nextMove.MoveType == "stun" {
		p.enemy.Stunned += p.nextMove.Value
		p.output.Add(fmt.Sprintf("%s uses %s! %s is stunned!", p.Name, p.nextMove.Name, p.enemy.Name))
		p.output.Render()
		time.Sleep(2 * time.Second)
		return
	}
	if p.nextMove.MoveType == "sacrifice" {
		p.enemy.Hp -= p.nextMove.Value
		p.Hp -= p.nextMove.SacValue

		p.output.Add(fmt.Sprintf("%s uses %s!", p.Name, p.nextMove.Name))
		p.output.Render()
		time.Sleep(2 * time.Second)
		return
	}

}

func (p player) showMoves() {
	for i := range p.Moves {
		if p.Moves[i].Uses == -1 {
			p.output.Add(fmt.Sprintf("[%d]  %s [%d]", i, p.Moves[i].Name, p.Moves[i].Value))
		} else {
			p.output.Add(fmt.Sprintf("[%d]  %s [%d] (Uses: %d)", i, p.Moves[i].Name, p.Moves[i].Value, p.Moves[i].Uses))
		}
		p.output.Render()
	}
	fmt.Println()
}

func (p player) MovePrompt() {
	p.output.Add(fmt.Sprintf("%s's move: ", p.Name))
	p.output.Add("")
	p.output.Render()
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
