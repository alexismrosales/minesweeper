package main

import (
	"client/connection"
	"client/game"
)

func main() {
	// Start connection
	conn := connection.ConnectToServer("localhost", "8080")
	// Close connection at the end of function
	defer conn.Close()

	board := game.BoardConstructor()
	board.StartGame(conn)
}
