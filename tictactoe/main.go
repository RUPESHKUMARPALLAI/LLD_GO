package main

import (
	"tictactoe/engine"
	"tictactoe/model"
)

func main() {
	game := engine.InitGame(3)
	game.AddPlayer("Rupesh", model.PieceTypeX)
	game.AddPlayer("Tinku", model.PieceTypeO)
	game.Play()
}
