package main

import (
	"fmt"
	"tic_tac_toe_game/game"
	gamexyz "tic_tac_toe_game/gameXyz"
	interfaceimpl "tic_tac_toe_game/interfaceImpl"
)

func main() {

	// g1, err := game.NewGame("Aniket", "Ankush", "X", "O")
	// if err != nil {
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	return
	// }
	// g.Play(0)
	// g.Play(2)
	// g.Play(8)
	// g.Play(10)
	// g.Play(8)
	// g.Play(6)
	// g.Play(4)
	var err error
	var g1 interfaceimpl.GameInterface

	g1, err = game.NewGame("Aniket", "Ankush", "X", "O")
	if err != nil {
		fmt.Println(err)
		return
	}
	g1.Play(3)
	g1.Play(0)
	g1.Play(1)
	g1.Play(6)

	fmt.Println("---------------------------------------------")
	var g2 interfaceimpl.GameInterface
	g2, err = gamexyz.NewGame("Vishav", "Brijesh", "@", "&")
	if err != nil {
		fmt.Println(err)
		return
	}
	g2.Play(2)
	g2.Play(4)
	g2.Play(1)
	g2.Play(0)
	g2.Play(8)

	g1 = g2

	g1.Printer()
	g2.Printer()

}
