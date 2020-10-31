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
}

func (p *player) doNextMove() {
	if p.Stunned > 0 {
		fmt.Printf("%s tries to use %s but they are STUNNED and cannot move!!", p.Name, p.nextMove.Name)
		p.Stunned -= 1
		time.Sleep(2 * time.Second)
		return
	}

	if p.nextMove.MoveType == "heal" {
		p.Hp += p.nextMove.Value
		fmt.Printf("%s uses %s and heals for %d", p.Name, p.nextMove.Name, p.nextMove.Value)
		time.Sleep(2 * time.Second)
		return
	}
	if p.nextMove.MoveType == "attack" {
		p.enemy.Hp -= p.nextMove.Value
		fmt.Printf("%s uses %s! %s takes %d damage.", p.Name, p.nextMove.Name, p.enemy.Name, p.nextMove.Value)
		time.Sleep(2 * time.Second)
		return
	}
	if p.nextMove.MoveType == "stun" {
		p.enemy.Stunned += p.nextMove.Value
		fmt.Printf("%s uses %s! %s is stunned!", p.Name, p.nextMove.Name, p.enemy.Name)
		time.Sleep(2 * time.Second)
		return
	}

}

func (p player) showMoves() {
	for i := range p.Moves {
		if p.Moves[i].Uses == -1 {
			fmt.Printf("[%d]  %s\n", i, p.Moves[i].Name)
		} else {
			fmt.Printf("[%d]  %s (Uses: %d)\n", i, p.Moves[i].Name, p.Moves[i].Uses)
		}
	}
	fmt.Println()
}

func (p player) MovePrompt() {
	fmt.Printf("%s's move: ", p.Name)
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
