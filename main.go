package main

import (
	"fmt"
	"bufio"
    "os"
    "strconv"
)

type move struct{
    name string
    movetype string
    value int
}

func (m move)execute(p *player) { 
    
}


type player struct {
	name    string
    health int
    moves []move
}


type game struct {
	p1   *player
	p2   *player
	turn int
}

func (p player)showOpts() { 
    fmt.Println(p.moves)
}


func main() {
	var players []*player
    healme := move{name: "healme", value: 3, movetype: "HEAL"}
    p1 := &player{name: "P1", health: 100, moves: []move{healme}}

	players = append(players, p1)

	for turn := 0; turn < 10; turn++ {
		for _, p := range players {
            p.showOpts() 
            cmd := GetInput()
            fmt.Println(cmd)
		}
	}
}

func GetInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	inputChar := input[0:1]
    val, _ := strconv.Atoi(inputChar)
    return val
}
