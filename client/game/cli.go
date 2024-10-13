package game

import "fmt"

var phrases = [...]string{
	"Selecciona la dificultad...\n1)Principiante\n2)Intermedio\n3)Experto",
	"Seleccione la acción a realizar:\n1)Buscar mina\n2)Colocar bandera",
	"Ingrese la coordenada del eje vertical (Numero)",
	"Ingrese la coordenada del eje horizontal (Letra)",
	"Felicidades, has ganado!",
	"Lo siento! Has perdido ...",
	"XXX Coordenadas incorrectas, porfavor verifiquelas y vuelva a introducirlas",
	"Se ha alcanzado el número máximo de banderas...",
}

var abecedary = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G"}

func askDifficulty() (int, int) {
	var h, w int
	// Set as default easy difficulty
	difficulty := 1
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
	return h, w
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
