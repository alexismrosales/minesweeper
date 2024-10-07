package connection

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

func ConnectToServer(address, port string) net.Conn {
	// Opening socket with TCP protocol
	conn, err := net.Dial("tcp", address+":"+port)
	// Handle error
	if err != nil {
		log.Fatal("Error conectandose al servidor: ", err)
	}
	fmt.Println("Conexión Exitosa...")
	return conn
}

func SendBoard(conn net.Conn, provider BoardProvider) {
	// Create a new encoder to serialize the object with the matrix
	encoder := gob.NewEncoder(conn)
	board := getBoardType(provider)
	// Encode the object
	err := encoder.Encode(board)
	if err != nil {
		log.Panic("Error enviando el objeto: ", err)
	}
}

func RecieveBoard(conn net.Conn) *Board {
	var board Board
	// Create a new decoder to deserialize the object
	decoder := gob.NewDecoder(conn)
	// Decoding the and save the objetct
	err := decoder.Decode(&board)

	if err != nil {
		log.Panic("Error reciviendo el objeto: ", err)
	}
	return &board
}
