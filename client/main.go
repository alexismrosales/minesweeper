package main

import (
	"client/connection"
	"client/game"
	"fmt"
)

func main() {
	var direction, port string
	fmt.Println("Escriba la dirección de la que desea conectarse: ")
	fmt.Scanln(&direction)
	fmt.Println("Escriba el puerto al que desee conectarse:")
	fmt.Scanln(&port)
	// Start connection
	conn := connection.ConnectToServer(direction, port)
	// Close connection at the end of function
	defer conn.Close()

	board := game.BoardConstructor()
	board.StartGame(conn)
}
