package model

type Player struct {
  Name string
  PlayingPiece PieceType
}

func NewPlayer(name string, piece PieceType) *Player {
  return &Player{
  Name: name,
  PlayingPiece: piece,
  }
}