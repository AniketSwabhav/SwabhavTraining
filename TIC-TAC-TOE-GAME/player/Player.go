package player

import "errors"

type Player struct {
	Name   string
	Symbol string
}

func NewPlayer(name, symbol string) (*Player, error) {
	if name == "" {
		return nil, errors.New("player name cannot be empty")
	}
	if symbol == "" {
		return nil, errors.New("symbol cannot empty")
	}

	return &Player{
		Name:   name,
		Symbol: symbol,
	}, nil
}
