package main

import (
	"fmt"
	"time"
)

func CharacterSelectMenu(player int, chars Characters, output *output) Elite {
	fmt.Printf("\n\nMeet the Elites!\n\n")

	fmt.Printf("[ # ]\tName\t\tMoves\n")
	fmt.Printf("=============================\n")

	for i, c := range chars.Elites {
		fmt.Printf("[ %d ]\t", i)
		fmt.Printf("%s", c.Name)
		for i, x := range c.Moves {
			switch i {
			case 0:
				if len(c.Name) > 8 {
					fmt.Printf("\t%s (%s)\n", x.Name, x.MoveType)
				} else {
					fmt.Printf("\t\t%s (%s)\n", x.Name, x.MoveType)
				}
			case 1:
				fmt.Printf("\t%d hp\t\t%s (%s)\n", c.Hp, x.Name, x.MoveType)
			default:
				fmt.Printf("\t\t\t%s (%s)\n", x.Name, x.MoveType)
			}
		}
		fmt.Println()
	}
	fmt.Printf("\n\nPlayer %d! Select your fighter: ", player)
	choice := GetInput()
	fmt.Printf("\n\nPlayer %d selects %s! A wise choice!\n", player, chars.Elites[choice].Name)
	time.Sleep(2 * time.Second)
	s.Output.Clear()
	return chars.Elites[choice]
}

func MainMenu(output *output) {
	output.Clear()
	delay := 50 * time.Millisecond

	for i := 0; i < 13; i++ {
		fmt.Printf("*")
		time.Sleep(delay)
	}
	time.Sleep(delay)
	fmt.Printf("\n*   ELITE   *\n")
	fmt.Printf("*           *\n")
	fmt.Printf("*  MONSTER  *\n")
	fmt.Printf("*           *\n")
	fmt.Printf("*   CARDS   *\n")
	time.Sleep(delay)
	for i := 0; i < 13; i++ {
		fmt.Printf("*")
		time.Sleep(delay)
	}
	time.Sleep(delay)
	fmt.Printf("\n\nPress enter to continue...")
	_ = GetInput()

}
