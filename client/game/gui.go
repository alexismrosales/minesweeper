package game

import (
	"client/connection"
	"fmt"
	"log"
	"net"
)

var phrases = [...]string{
	"Selecciona la dificultad...\n1)Principiante\n2)Intermedio\n3)Experto",
	"Seleccione la acción a realizar:\n1)Buscar mina\n2)Colocar bandera",
	"Ingrese la coordenada del eje vertical (Numero)",
	"Ingrese la coordenada del eje horizontal (Letra)",
	"Lo siento! Has perdido ...",
	"Felicidades, has ganado!",
	"XXX Coordenadas incorrectas, porfavor verifiquelas y vuelva a introducirlas",
}

var abecedary = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G"}

func BoardConstructor() *Board {
	return &Board{}
}

func (board *Board) StartGame(conn net.Conn) {
	difficulty := 1
	var h, w int
	// Ask for difficulty
	fmt.Println(phrases[0])
	fmt.Scanln(&difficulty)
	// Generate board size
	switch difficulty {
	case 1:
		h, w = 9, 9
	case 2:
		h, w = 16, 16
	case 3:
		h, w = 16, 32
	default:
		h, w = 9, 9
	}
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

		// Save the selected value
		board.Values[board.X][board.Y] = value
		if value == '$' {
			continue
		}
		// Show selected value
		board.printBoard()

		// Send and recieve of the object
		connection.SendBoard(conn, board)
		log.Println("Se envió el objeto correctamente...")
		boardUpdated := connection.RecieveBoard(conn)
		log.Println("El objeto se recibio correctamente")

		// Board modified by the server
		board = getBoardType(boardUpdated)

		printStatus(board.Status)
	}
}

func printStatus(status int) {
	switch status {
	case 0:
		return
	case 1:
		fmt.Println(phrases[4])
	case 2:
		fmt.Println(phrases[5])
	}
}
func askOption(coordenateX *int, coordenateY *string) rune {
	var option int
	// Show options
	fmt.Println(phrases[1])
	fmt.Scanln(&option)
	fmt.Println(phrases[3])
	fmt.Scanln(coordenateY)
	fmt.Println(phrases[2])
	fmt.Scanln(coordenateX)
	fmt.Print("\033[H\033[2J")

	if option == 2 {
		return '$'
	}
	return '*'
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

func (board Board) printBoard() {
	h, w := board.H, board.W
	values := board.Values
	// Iterate over matrix and print values
	for i := 0; i <= h; i++ {
		for j := 0; j <= w; j++ {
			// First row
			if i == 0 {
				// First cell
				if j != 0 {
					fmt.Print(abecedary[j-1])
				}
			} else if j == 0 {
				fmt.Print(i)
			} else {
				fmt.Print(string(values[i-1][j-1]))
			}
			fmt.Print("\t")
		}
		fmt.Print("\n\n")
	}
}
