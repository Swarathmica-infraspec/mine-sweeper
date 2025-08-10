package minesweeper

type Box struct {
    x, y  int
    state rune
}

func NewBox(x, y int, state rune) *Box {
    return &Box{x: x, y: y, state: state}
}
