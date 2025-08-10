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
