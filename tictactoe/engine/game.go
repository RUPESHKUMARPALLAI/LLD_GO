package engine

import (
    "tictactoe/model"
)

type Game struct {
    Board *model.Board
    Players []*model.Player
    MoveCount int
    CurrentPlayer int
    Winner *model.Player
}

func InitGame (boardSize int) *Game {
    return &Game{
    Board : model.InitBoard(boardSize),
    Players : []*model.Player{},
    MoveCount : 0,
    CurrentPlayer : 0,
    Winner : nil,
    }
}

func (g *Game) AddPlayer (name string, pieceType model.PieceType) {
    player := model.NewPlayer(name, pieceType)
    g.Players = append(g.Players, player)
}