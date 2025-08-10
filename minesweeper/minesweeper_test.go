package minesweeper

import (
	"errors"
	"testing"
)

func TestNewBox(t *testing.T) {
	b, _ := NewBox(1, 2, 'E')
	if b.x != 1 || b.y != 2 || b.state != 'E' {
		t.Errorf("minesweeper: NewBox: Box not created properly, got %+v", b)
	}
}

func TestNewBoxCannotBeCreatedWithNegativeCoordinatesX(t *testing.T) {
	_, err := NewBox(-1, 2, 'E')
	if !errors.Is(err, errors.ErrUnsupported) {
		t.Errorf("minesweeper: NewBox: Cannot create Box with negative coordinates")
	}
}

func TestNewBoxCannotBeCreatedWithNegativeCoordinatesY(t *testing.T) {
	_, err := NewBox(1, -2, 'E')
	if !errors.Is(err, errors.ErrUnsupported) {
		t.Errorf("minesweeper: NewBox: Cannot create Box with negative coordinates")
	}
}

func TestNewBoxCannotBeCreatedWithStateOtherThanMOrE(t *testing.T) {
	_, err := NewBox(1, -2, 'N')
	if !errors.Is(err, errors.ErrUnsupported) {
		t.Errorf("minesweeper: NewBox: Cannot create Box with state other than M or E")
	}
}

var boardState = [][]rune{
	{'E', 'E'},
	{'E', 'E'},
}

func TestNewBoard(t *testing.T) {

	_, err := NewBoard(2, 2, boardState)

	if err != nil {
		t.Errorf("minesweeper: NewBoard: expected no error, got %v", err)
	}

}

func TestNewBoardannotBeCreatedWithNegativeXDimension(t *testing.T) {
	_, err := NewBoard(-2, 2, boardState)
	if err == nil {
		t.Errorf("minesweeper: NewBoard: cannot create board with dimensions<1")
	}

}

func TestNewBoardannotBeCreatedWithNegativeYDimension(t *testing.T) {
	_, err := NewBoard(2, -2, boardState)
	if err == nil {
		t.Errorf("minesweeper: NewBoard: cannot create board with dimensions<1")
	}

}

func TestNeighbours(t *testing.T) {
	b, _ := NewBoard(3, 3, [][]rune{
		{'E', 'E', 'E'},
		{'E', 'M', 'E'},
		{'E', 'E', 'E'},
	})

	neighbours := b.Neighbours(1, 1)

	expectedCoords := [][2]int{
		{0, 0}, {1, 0}, {2, 0},
		{0, 1}, {2, 1},
		{0, 2}, {1, 2}, {2, 2},
	}

	if len(neighbours) != len(expectedCoords) {
		t.Errorf("minesweeper: Neighbous: expected %d neighbours, got %d", len(expectedCoords), len(neighbours))
	}

	for _, coord := range expectedCoords {
		found := false
		for _, nb := range neighbours {
			if nb.x == coord[0] && nb.y == coord[1] {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("minesweeper: Neighbous: expected neighbour at (%d,%d) not found", coord[0], coord[1])
		}
	}
}
