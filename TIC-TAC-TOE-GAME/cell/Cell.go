package cell

type Cell struct {
	Mark string
}

func NewCell() *Cell {
	return &Cell{
		Mark: " ",
	}
}
