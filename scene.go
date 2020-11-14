package main

import (
	"fmt"
	"time"
)

func CharacterSelectMenu(player int, chars Characters, output *output) Elite {
	output.Clear()
	output.Add(fmt.Sprintf("[ # ]\tName\t\tMoves\n"))
	output.Add(fmt.Sprintf("=============================\n"))
	output.Render()
	for i, c := range chars.Elites {
		output.Add(fmt.Sprintf("[ %d ]\t", i))
		output.Add(fmt.Sprintf("%s", c.Name))
		for i, x := range c.Moves {
			switch i {
			case 0:
				if len(c.Name) > 8 {
					output.Add(fmt.Sprintf("\t%s (%s)\n", x.Name, x.MoveType))
				} else {
					output.Add(fmt.Sprintf("\t\t%s (%s)\n", x.Name, x.MoveType))
				}
			case 1:
				output.Add(fmt.Sprintf("\t%d hp\t\t%s (%s)\n", c.Hp, x.Name, x.MoveType))
			default:
				output.Add(fmt.Sprintf("\t\t\t%s (%s)\n", x.Name, x.MoveType))
			}
		}
		output.Add("\n")
		output.Render()
	}
	output.Add(fmt.Sprintf("\n\nPlayer %d! Select your fighter: ", player))
	output.Render()
	choice := GetInput()
	output.Add(fmt.Sprintf("\n\nPlayer %d selects %s! A wise choice!\n", player, chars.Elites[choice].Name))
	output.Render()
	time.Sleep(2 * time.Second)
	output.Clear()
	return chars.Elites[choice]
}

func MainMenu(output *output) {
	output.Clear()
	delay := 50 * time.Millisecond

	for i := 0; i < 13; i++ {
		output.Add(fmt.Sprintf("*"))
		time.Sleep(delay)
	}
	time.Sleep(delay)
	output.Add(fmt.Sprintf("\n*   ELITE   *\n"))
	output.Add(fmt.Sprintf("*           *\n"))
	output.Add(fmt.Sprintf("*  MONSTER  *\n"))
	output.Add(fmt.Sprintf("*           *\n"))
	output.Add(fmt.Sprintf("*   CARDS   *\n"))
	time.Sleep(delay)
	for i := 0; i < 13; i++ {
		output.Add("*")
		output.Render()
		time.Sleep(delay)
	}
	time.Sleep(delay)
	output.Add(fmt.Sprintf("\n\nPress enter to continue..."))
	output.Render()
	_ = GetInput()

}
