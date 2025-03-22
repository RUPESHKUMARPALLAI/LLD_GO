package model

import "fmt"

type Board struct {
	Size  int
	Cells [][]PieceType
}

func InitBoard(size int) *Board {
	grid := make([][]PieceType, size)
	for i := range grid {
		grid[i] = make([]PieceType, size)
		for j := range grid[i] {
			grid[i][j] = PieceEmpty
		}
	}
	return &Board{Size: size, Cells: grid}
}

func (b *Board) SetCell(x int, y int, piece PieceType) {
	if b.Cells[x][y] == PieceEmpty {
		b.Cells[x][y] = piece
	} else {
		fmt.Println("Cell is not empty")
	}
}

func (b *Board) PrintBoard() {
	for _, row := range b.Cells {
		for _, cell := range row {
			fmt.Print(cell + "  ")
		}
		print("\n")
	}
}
