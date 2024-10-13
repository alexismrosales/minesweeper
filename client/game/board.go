package game

import "client/connection"

type Board struct {
	X, Y             int
	H, W             int
	Status           int
	Values           [][]rune
	GameValues       [][]rune
	MinesCoordinates map[[2]int]struct{}
}

func (b *Board) GetValues() *[][]rune {
	return &b.Values
}

func (b *Board) GetGameValues() *[][]rune {
	return &b.GameValues
}

func (b *Board) GetCoordinates() (int, int) {
	return b.X, b.Y
}

func (b *Board) GetDimensions() (int, int) {
	return b.H, b.W
}

func (b *Board) GetStatus() int {
	return b.Status
}

func (b *Board) GetMinesCoordinates() map[[2]int]struct{} {
	return b.MinesCoordinates
}

func getBoardType(board *connection.Board) *Board {
	h, w := board.GetDimensions()
	x, y := board.GetCoordinates()
	return &Board{
		H:                h,
		W:                w,
		X:                x,
		Y:                y,
		Values:           *board.GetValues(),
		Status:           board.GetStatus(),
		GameValues:       *board.GetGameValues(),
		MinesCoordinates: board.GetMinesCoordinates(),
	}
}
