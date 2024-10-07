package extras

import (
	"log"
	"os"
)

func SaveRecord(time int) {
	// Create or append new file to write text
	file, err := os.OpenFile("records.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Println("Error al abrir o crear el archivo: ", err)
		return
	}
	// Close file
	defer file.Close()
	// Writing time in file
	_, err = file.WriteString("Tiempo de " + string(time) + "segundos. \n")
	if err != nil {
		log.Println("Error al escribir el tiempo en el archivo records.txt :", err)
		return
	}

	log.Println("Récord agregado correctamente...")
}
