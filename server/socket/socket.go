package socket

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"server/extras"
	"server/game"
	"time"
)

func InitializeServer(port string) {
	// Create socket connection with TCP protocol
	listener, err := net.Listen("tcp", ":"+port)
	// Handle error
	if err != nil {
		log.Fatal("Error al crear la conexion: ", err)
	}
	// Close listener at the end of the function
	defer listener.Close()
	/*
		  // Se empezarán a recibir una por una todas las peticiones que sean mandados al
			// servidor, una vez lograda la conexión empezará la comunicación de datos con
			// el cliente.
	*/
	// Accept all connections
	for {
		// Block stream until a client connects
		conn, err := listener.Accept()
		// Handle connection error
		if err != nil {
			log.Panic("Error aceptando la conexión", err)
		}
		//    Start a new match with the player
		// Start of timer
		start := time.Now()
		// Start of the game
		handleConnection(conn)
		// End of timer
		duration := time.Since(start)
		fmt.Println("La duración de la partida fue de ", int(duration.Seconds())%60, "segundos.")
		extras.SaveRecord(int(duration))
	}

}

func handleConnection(conn net.Conn) {
	// Close connection
	defer conn.Close()
	fmt.Println("Cliente conectado: ", conn.LocalAddr().String())
	/*
	   // Se recibe el objeto (tablero) del  cliente y se transforma
	   // al tipo de dato necesario para poder manipular el objeto,
	   // una vez realizada la lógica del tablero se enviará devuelta
	   // al cliente
	*/
	// Listen responses and send responses
	for {
		board, err := RecieveBoard(conn)
		if err != nil {
			if err == io.EOF {
				log.Println("Conexión cerrada por el cliente")
				break
			}
			log.Panic("Error recibiendo el objeto: ", err)
		}
		bboard := game.ModifyBoard(board)
		board = getBoardType(bboard)
		SendBoard(conn, *board)
	}
}

func SendBoard(conn net.Conn, board Board) {
	// Create a new encoder to serialize the object with the matrix
	encoder := gob.NewEncoder(conn)
	// Encode the object
	err := encoder.Encode(board)
	if err != nil {
		log.Panic("Error enviando el objeto: ", err)
	}
}

func RecieveBoard(conn net.Conn) (*Board, error) {
	var board Board
	// Create a new decoder to deserialize the object
	decoder := gob.NewDecoder(conn)
	// Decoding the and save the objetct
	err := decoder.Decode(&board)

	if err != nil {
		return &board, err
	}
	return &board, nil
}
