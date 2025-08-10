package minesweeper

import "testing"

func TestNewBox(t *testing.T) {
	b := NewBox(1, 2, 'E')
	if b.x != 1 || b.y != 2 || b.state != 'E' {
		t.Errorf("Box not created properly, got %+v", b)
	}
}
