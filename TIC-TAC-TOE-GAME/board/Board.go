package board

import (
	"errors"
	"fmt"
	"tic_tac_toe_game/cell"
)

type Board struct {
	Cells [9]*cell.Cell
}

func NewBoard() *Board {

	var cellArray [9]*cell.Cell

	for i := 0; i < len(cellArray); i++ {
		cellArray[i] = cell.NewCell()
	}

	return &Board{
		Cells: cellArray,
	}
}

func (b *Board) IsCellEmpty(pos int) bool {
	if pos < 0 || pos >= len(b.Cells) {
		return false
	}
	return b.Cells[pos].Mark == " "
}

func (b *Board) MarkCell(pos int, symbol string) error {
	if pos < 0 || pos >= len(b.Cells) {
		return errors.New("invalid cell index")
	}
	if b.Cells[pos].Mark != " " {
		return errors.New("cell is already marked")
	}
	b.Cells[pos].Mark = symbol
	return nil
}

func (b *Board) CheckWin(symbol string) bool {
	winPatterns := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}
	for _, pattern := range winPatterns {
		if b.Cells[pattern[0]].Mark == symbol &&
			b.Cells[pattern[1]].Mark == symbol &&
			b.Cells[pattern[2]].Mark == symbol {
			return true
		}
	}
	return false
}

func (b *Board) CheckDraw() bool {
	for _, c := range b.Cells {
		if c.Mark == " " {
			return false
		}
	}
	return true
}

func (b *Board) Reset() {
	for _, c := range b.Cells {
		c.Mark = " "
	}
}

func (b *Board) PrintBoard() {
	for i := 0; i < 9; i += 3 {
		fmt.Printf(" %s | %s | %s \n", b.Cells[i].Mark, b.Cells[i+1].Mark, b.Cells[i+2].Mark)
		if i < 6 {
			fmt.Println("---|---|---")
		}
	}
}
