package main

import "fmt"

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
