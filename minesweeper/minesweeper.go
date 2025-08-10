package minesweeper

import "errors"

type Box struct {
	x, y  int
	state rune
}

func NewBox(x, y int, state rune) (*Box, error) {
	if x < 0 || y < 0 || state != 'M' && state != 'E' {
		return nil, errors.ErrUnsupported
	}
	return &Box{x: x, y: y, state: state}, nil
}
