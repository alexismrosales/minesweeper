package socket

import "server/game"

type Board struct {
	X, Y             int      // Coordinates of the player last turn
	H, W             int      // Size of height and weight of the matrix
	Values           [][]rune // Matrix of values
	GameValues       [][]rune
	Status           int
	MinesCoordinates map[[2]int]struct{}
}

func (b *Board) GetValues() *[][]rune {
	return &b.Values
}

func (b *Board) GetGameValues() *[][]rune {
	return &b.GameValues
}

// Métodos adicionales para obtener coordenadas y dimensiones
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

func getBoardType(board *game.Board) *Board {
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
