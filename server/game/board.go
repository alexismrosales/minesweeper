package game

type Board struct {
	X, Y       int      // Coordinates of the player last turn
	H, W       int      // Size of height and weight of the matrix
	Values     [][]rune // Matrix of values
	GameValues [][]rune
	Status     int
}

type BoardProvider interface {
	GetValues() *[][]rune
	GetCoordinates() (int, int)
	GetDimensions() (int, int)
	GetStatus() int
	GetGameValues() *[][]rune
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

func getBoardType(board BoardProvider) *Board {
	h, w := board.GetDimensions()
	x, y := board.GetCoordinates()
	return &Board{
		H:          h,
		W:          w,
		X:          x,
		Y:          y,
		Values:     *board.GetValues(),
		Status:     board.GetStatus(),
		GameValues: *board.GetGameValues(),
	}
}

func getTotalMines(board *Board) int {
	w := board.W
	switch w {
	case 9:
		// Easy
		return 10
	case 16:
		// Medium
		return 40
	case 30:
		// Hard
		return 99
	}
	return 0
}
