package minesweeper

import "errors"

type Box struct {
	x, y  int
	state rune
}

type Board struct {
	width, height int
	Boxes         [][]*Box
}

func NewBox(x, y int, state rune) (*Box, error) {
	if x < 0 || y < 0 || state != 'M' && state != 'E' {
		return nil, errors.ErrUnsupported
	}
	return &Box{x: x, y: y, state: state}, nil
}

func NewBoard(width, height int, boardState [][]rune) (*Board, error) {
	if width < 1 || height < 1 {
		return nil, errors.ErrUnsupported
	}
	boxes := make([][]*Box, height)
	for y := 0; y < height; y++ {
		boxes[y] = make([]*Box, width)
		for x := 0; x < width; x++ {

			b, err := NewBox(x, y, boardState[y][x])
			if err != nil {
				return nil, errors.New(err.Error())
			}

			boxes[y][x] = b
		}
	}
	return &Board{Boxes: boxes, width: width, height: height}, nil
}

func (b Board) Neighbours(x, y int) []*Box {
	var neighbours []*Box
	directions := [][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	for _, d := range directions {
		nx := x + d[0]
		ny := y + d[1]
		if nx >= 0 && nx < b.width && ny >= 0 && ny < b.height {
			neighbours = append(neighbours, b.Boxes[ny][nx])
		}

	}

	return neighbours
}

func (b *Board) CountAdjacentMines(x, y int) int {
	count := 0
	for _, box := range b.Neighbours(x, y) {
		if box.state == 'M' {
			count++
		}
	}
	return count
}
