package engine

import (
	"fmt"
	"tictactoe/model"
)

func (g *Game) Play() {
	currentPlayer := g.Players[g.CurrentPlayer]
	for {
		row, col := 0, 0
		for {
			fmt.Printf("%s's turn (%s). Enter row and column (0-based index): ", currentPlayer.Name, currentPlayer.PlayingPiece)
			fmt.Scan(&row, &col)
			if row < 0 || row >= g.Board.Size || col < 0 || col >= g.Board.Size {
				fmt.Println("❌ Invalid move: Out of bounds.")
				continue
			}
			if g.Board.Cells[row][col] != model.PieceEmpty {
				fmt.Println("❌ Invalid move: Cell already taken.")
				continue
			}
			break

		}

		g.Board.Cells[row][col] = currentPlayer.PlayingPiece
		g.MoveCount++
		g.Board.PrintBoard()

		if CheckWinner(g.Board, row, col, currentPlayer.PlayingPiece) {
			g.Winner = currentPlayer
			fmt.Printf("🎉 Winner: %s!\n", currentPlayer.Name)
			return
		}

		if g.MoveCount == g.Board.Size*g.Board.Size {
			fmt.Printf("Game is Draw!\n")
			return
		}
		g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
		currentPlayer = g.Players[g.CurrentPlayer]
	}
}

func CheckWinner(board *model.Board, row, col int, piece model.PieceType) bool {
	size := board.Size
	win := true
	for i := 0; i < size; i++ {
		if board.Cells[row][i] != piece {
			win = false
			break
		}
	}
	if win {
		return true
	}

	win = true

	for i := 0; i < size; i++ {
		if board.Cells[i][col] != piece {
			win = false
			break
		}
	}
	if win {
		return true
	}

	if row == col {
		win = true
		for i := 0; i < size; i++ {
			if board.Cells[i][i] != piece {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	if row+col == size-1 {
		win = true
		for i := 0; i < size; i++ {
			if board.Cells[i][size-i-1] != piece {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	return win
}
