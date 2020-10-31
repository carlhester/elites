package main

import (
	"fmt"
	"time"
)

func CharacterSelectMenu(player int, chars Characters) Elite {
	fmt.Printf("\n\nMeet the Elites!\n\n")

	for i, c := range chars.Elites {
		fmt.Printf("[%d]  ", i)
		fmt.Printf("Name: %s\t", c.Name)
		fmt.Printf("HP: %d\t", c.Hp)
		fmt.Printf("Moves: ")
		for _, x := range c.Moves {
			fmt.Printf("%s (%s)\t", x.Name, x.MoveType)
		}
		fmt.Println()
	}
	fmt.Printf("\n\nPlayer %d! Select your fighter: ", player)
	choice := GetInput()
	fmt.Printf("\n\n%s! Wise choice!\n", chars.Elites[choice].Name)
	time.Sleep(1 * time.Second)
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
