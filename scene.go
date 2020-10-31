package main

import (
	"fmt"
	"time"
)

func CharacterSelectMenu(player int, chars Characters) Elite {
	fmt.Printf("\n\nMeet the Elites!\n\n")

	fmt.Printf("[ # ]\tName\tHp\tMoves\n")
	fmt.Printf("=============================\n")

	for i, c := range chars.Elites {
		fmt.Printf("[ %d ]\t", i)
		fmt.Printf("%s\t", c.Name)
		fmt.Printf("%d\t", c.Hp)
		for _, x := range c.Moves {
			fmt.Printf("%s (%s)\n\t\t\t", x.Name, x.MoveType)
		}
		fmt.Println()
	}
	fmt.Printf("\n\nPlayer %d! Select your fighter: ", player)
	choice := GetInput()
	fmt.Printf("\n\nPlayer %d selects %s! A wise choice!\n", player, chars.Elites[choice].Name)
	time.Sleep(2 * time.Second)
	Clear()
	return chars.Elites[choice]
}

func MainMenu() {
	Clear()
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
