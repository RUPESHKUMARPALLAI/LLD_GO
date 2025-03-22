package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard(t *testing.T) {
	size := 3
	b := InitBoard(3)

	assert.NotNil(t, b, "Shouldnt be nil")

	assert.Equal(t, size, b.Size, "size should match")

	assert.Equal(t, b.Cells[0][0], PieceEmpty, "Initialize Piece is _")

	b.SetCell(0, 0, PieceTypeO)
	
	assert.Equal(t, b.Cells[0][0], PieceTypeO, "Asserting setPiece")

}
