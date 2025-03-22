package engine

import (
	"testing"
	"tictactoe/model"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	g := InitGame(3)
	assert.NotNil(t, g, "Game object should not be nil")
	assert.Equal(t, g.Board.Size, 3, "Asserting Game BoardSize")
	assert.Equal(t, 0, len(g.Players), "Players slice should be empty initially")
	assert.Nil(t, g.Winner, "Winner is Nil")
	g.AddPlayer("Rupesh", model.PieceTypeO)
	assert.Equal(t, g.Players[0].Name, "Rupesh", "Player Name Assertion")
}