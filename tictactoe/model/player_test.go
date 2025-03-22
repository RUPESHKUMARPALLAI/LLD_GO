package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer(t *testing.T) {
	p := NewPlayer("Rupesh", PieceTypeO)
	assert.Equal(t, p.Name, "Rupesh", "Name Assertion")
	assert.Equal(t, p.PlayingPiece, PieceTypeO, "PieceType Assertion")
}
