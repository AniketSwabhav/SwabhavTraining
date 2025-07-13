package game

import (
	"fmt"
	"tic_tac_toe_game/board"
	"tic_tac_toe_game/player"
	"tic_tac_toe_game/util"
)

type Game struct {
	Players    [2]player.Player
	Board      board.Board
	Turn       int
	Winner     player.Player
	IsDraw     bool
	IsGameOver bool
}

func NewGame(player1, player2 string, player1Symbol, player2Symbol string) (*Game, error) {

	p1, err := player.NewPlayer(player1, player1Symbol)
	if err != nil {
		return nil, err
	}
	p2, err := player.NewPlayer(player2, player2Symbol)
	if err != nil {
		return nil, err
	}

	board := board.NewBoard()

	return &Game{
		Players:    [2]player.Player{*p1, *p2},
		Board:      *board,
		Turn:       0,
		IsDraw:     false,
		IsGameOver: false,
	}, nil
}

func (g *Game) Play(position int) {

	defer util.HandlePanic()

	if g.IsGameOver {
		panic("game is already over")
	}

	currentPlayer := g.Players[g.Turn%2]

	err := g.Board.MarkCell(position, currentPlayer.Symbol)
	if err != nil {
		panic(err)
	}

	g.Board.PrintBoard()

	if g.Board.CheckWin(currentPlayer.Symbol) {
		g.IsGameOver = true
		g.Winner = currentPlayer
		fmt.Printf("Game Over! Winner is %s (symbol '%s')\n", g.Winner.Name, g.Winner.Symbol)
	}

	if g.Board.CheckDraw() {
		g.IsGameOver = true
		g.IsDraw = true
		fmt.Println("Game Over! It's a draw.")
	}

	g.Turn++
}

func (g *Game) Reset() {
	g.Board.Reset()
	g.Turn = 0
	g.IsGameOver = false
	g.IsDraw = false
	g.Winner = player.Player{}
}

//surprise
//insight
//pride
//gratitude
