package engine

import (
	"testing"
	"tictactoe/model"

	"github.com/stretchr/testify/assert"
)

func TestWinConditionRow(t *testing.T) {
	board := model.InitBoard(3)

	// Simulate row win for PieceTypeX
	board.Cells[0][0] = model.PieceTypeX
	board.Cells[0][1] = model.PieceTypeX
	board.Cells[0][2] = model.PieceTypeX

	win := CheckWinner(board, 0, 2, model.PieceTypeX)
	assert.True(t, win)
}

func TestDrawCondition(t *testing.T) {
	game := InitGame(3)
	game.AddPlayer("X", model.PieceTypeX)
	game.AddPlayer("O", model.PieceTypeO)

	// Fill board with no winner
	board := [][]model.PieceType{
		{model.PieceTypeX, model.PieceTypeO, model.PieceTypeX},
		{model.PieceTypeX, model.PieceTypeO, model.PieceTypeO},
		{model.PieceTypeO, model.PieceTypeX, model.PieceTypeX},
	}
	game.Board.Cells = board
	game.MoveCount = 9

	// No winner
	assert.False(t, CheckWinner(game.Board, 2, 2, model.PieceTypeX))
	assert.Equal(t, 9, game.MoveCount)
}



