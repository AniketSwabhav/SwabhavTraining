package main

import (
	"fmt"
	"tic_tac_toe_game/game"
)

func main() {

	g, err := game.NewGame("Aniket", "Ankush", "X", "O")
	if err != nil {
		if err != nil {
			fmt.Println(err)
			return
		}

		return
	}
	g.Play(0)
	g.Play(2)
	g.Play(8)
	g.Play(10)
	g.Play(8)
	g.Play(6)
	g.Play(4)
}
