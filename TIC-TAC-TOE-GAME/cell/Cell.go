package cell

import "errors"

type Cell struct {
	Mark string
}

func NewCell() *Cell {
	return &Cell{
		Mark: " ",
	}
}

func (c *Cell) SetMark(mark string) error {
	if c.Mark != " " {
		return errors.New("Cell is already marked")
	}
	c.Mark = mark
	return nil
}

func (c *Cell) Clear() error {
	if c.Mark != " " {
		c.Mark = " "
	} else {
		return errors.New("Cell is already empty")
	}
	return nil
}
