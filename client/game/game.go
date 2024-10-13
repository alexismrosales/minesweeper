package game

import (
	"client/connection"
	"fmt"
	"log"
	"net"
)

func BoardConstructor() *Board {
	return &Board{}
}

func (board *Board) StartGame(conn net.Conn) {
	h, w := askDifficulty()
	// Add values to the struct
	board.H = h
	board.W = w
	board.Values = initializeValues(h, w)
	board.loopBoard(conn)
}

func initializeValues(h, w int) [][]rune {
	// Creating an empty matrix
	values := make([][]rune, h)
	for i := range values {
		values[i] = make([]rune, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			values[i][j] = '-'
		}
	}
	return values
}

func (board *Board) loopBoard(conn net.Conn) {
	coordenateX := 0
	coordenateY := ""
	flagCounter := 0
	h, w := board.H, board.W
	// gameCondition == 0 | User still in game;
	// gameCondition == 1 | User wins;
	// gameCondition == 2 | User lost;
	// Show board and info
	for board.Status == 0 {
		board.printBoard()
		value := askOption(&coordenateX, &coordenateY)
		board.X, board.Y = saveValues(coordenateX, coordenateY, h, w)
		// In case of one or more value is invalid
		if board.X == -1 || board.Y == -1 {
			fmt.Println(phrases[6])
			continue
		}
		if value == '$' {
			flagCounter++
			board.Values[board.X][board.Y] = value
			// Show selected value
			board.printBoard()
			continue
		}
		// Save the selected value
		board.Values[board.X][board.Y] = value
		// Show selected value
		board.printBoard()

		// Send and recieve of the object
		connection.SendBoard(conn, board)
		log.Println("Se envió el objeto correctamente...")
		boardUpdated := connection.RecieveBoard(conn)
		log.Println("El objeto se recibio correctamente")

		// Board modified by the server
		board = getBoardType(boardUpdated)
	}

	board.printBoard()
	fmt.Println(board.Status)
	printStatus(board.Status)
}

func saveValues(valX int, valY string, h, w int) (int, int) {
	indexY := -1
	indexX := valX - 1
	// Save value for Y axis or abecedary coordenate
	for i := range abecedary {
		if abecedary[i] == valY {
			indexY = i
		}
	}
	// Save value in the indicated coordinate
	if indexX+1 > h || indexX+1 < 0 {
		indexX = -1
	}
	if indexY > w || indexY < 0 {
		indexY = -1
	}
	return indexX, indexY
}
