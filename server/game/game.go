package game

import (
	"fmt"
	"math/rand"
)

func ModifyBoard(b BoardProvider) *Board {
	userBoard := getBoardType(b)
	if userBoard.Status != 0 {
		return userBoard
	}
	// First time
	h, w := userBoard.H, userBoard.W
	if len(userBoard.GameValues) == 0 {
		mines := getTotalMines(userBoard)
		userBoard.GameValues = initializeValues(h, w)
		generateBoard(h, w, mines, userBoard.GameValues)
	}
	printBoard(h, w, userBoard.GameValues)
	fmt.Println("-----\n")
	userBoard.Status = compareBoards(userBoard.X, userBoard.Y, userBoard.Values, userBoard.GameValues)
	playerWins(userBoard.Values, userBoard.GameValues)
	return userBoard
}

func compareBoards(userCoordinateX, userCoordinateY int, userValues [][]rune, boardValues [][]rune) int {
	directions := [4][2]int{
		{-1, 0}, // left
		{1, 0},  // right
		{0, 1},  // up
		{0, -1}, // down
	}
	/*
	  // Hay 3 casos posibles, al momento que el usuario eliga una coordenada:
	  // 1. Que se haya seleccionado una mina
	  // 2. Que haya seleccionado una casilla con numero
	  // 3. Que haya seleccionado una casilla vacia
	*/

	// First case
	if boardValues[userCoordinateX][userCoordinateY] == '#' {
		// Save the mine symbol on the user board
		userValues[userCoordinateX][userCoordinateX] = '#'
		return 2
	}
	// Second case
	if boardValues[userCoordinateX][userCoordinateY] != '-' {
		userValues[userCoordinateX][userCoordinateY] = boardValues[userCoordinateX][userCoordinateY]
	} else {
		// Third case
		recursiveReveal(userCoordinateX, userCoordinateY, userValues, boardValues, directions)
	}

	return 0
}

func recursiveReveal(userCoordinateX, userCoordinateY int, userValues [][]rune, boardValues [][]rune, directions [4][2]int) {
	// Base case
	if boardValues[userCoordinateX][userCoordinateY] != '-' {
		if boardValues[userCoordinateX][userCoordinateY] != '#' {
			userValues[userCoordinateX][userCoordinateY] = boardValues[userCoordinateX][userCoordinateY]
		}
		return
	}
	if userValues[userCoordinateX][userCoordinateY] == '/' {
		return
	}
	userValues[userCoordinateX][userCoordinateY] = '/'
	// Checking in every direction if there is a cell to reveal
	for _, direction := range directions {
		newX := userCoordinateX + direction[0]
		newY := userCoordinateY + direction[1]

		if newX >= 0 && newX < len(boardValues) && newY >= 0 && newY < len(boardValues[0]) {
			recursiveReveal(newX, newY, userValues, boardValues, directions)
		}
	}
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

func generateBoard(h, w, totalMines int, values [][]rune) {

	// Map of values with randomCoordinates
	minesPosition := make(map[[2]int]struct{})
	for i := 0; i < totalMines; i++ {
		//
		randomX := rand.Intn(w)
		randomY := rand.Intn(h)
		randomCoordinate := [2]int{randomX, randomY}
		// If the coordinate do not exists, is saved on the map
		if _, exist := minesPosition[randomCoordinate]; !exist {
			minesPosition[randomCoordinate] = struct{}{}
		} else {
			// Else try again
			i--
		}
	}
	generateNumbers(h, w, minesPosition, values)
	generateMines(minesPosition, values)
}

func generateNumbers(h, w int, minesPosition map[[2]int]struct{}, values [][]rune) {
	directions := [8][2]int{
		{-1, 0},  // left
		{1, 0},   // right
		{0, 1},   // up
		{0, -1},  // down
		{1, 1},   // up-right
		{-1, 1},  // up-left
		{-1, -1}, // down-left
		{1, -1},  // down-right
	}
	for minePosition := range minesPosition {
		x, y := minePosition[0], minePosition[1]
		// Put mines
		for _, direction := range directions {
			nx, ny := x+direction[0], y+direction[1]
			// Verify limits of the board
			if nx >= 0 && nx < w && ny >= 0 && ny < h {
				// Check if there is not a mine
				if values[nx][ny] != '#' {
					if values[nx][ny] == '-' {
						values[nx][ny] = '1'
					} else {
						// Converting rune to int an sum one and then save the new value
						newValue := int(values[nx][ny]-'0') + 1
						values[nx][ny] = rune(newValue + '0')
					}
				}
			}
		}
	}
}

func generateMines(minesPosition map[[2]int]struct{}, values [][]rune) {
	for minePosition := range minesPosition {
		x, y := minePosition[0], minePosition[1]
		values[x][y] = '#'
	}
}

func printBoard(h, w int, values [][]rune) {
	// Iterate over matrix and print values
	for i := 0; i <= h; i++ {
		for j := 0; j <= w; j++ {
			// First row
			if i == 0 {
				// First cell
				if j != 0 {
					fmt.Print(-1)
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

func playerWins(userValues [][]rune, boardValues [][]rune) bool {
	for i := 0; i < len(userValues); i++ {
		for j := 0; j < len(userValues[i]); j++ {
			if userValues[i][j] != boardValues[i][j] && userValues[i][j] != '-' {
				return false
			}
		}
	}
	return true
}
