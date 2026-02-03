package main

import (
	"log"
	"os"
)

func funciones_log() {
	log.SetPrefix("main():") // prefijo que tendra el log, aparece a la izquierda
	log.Printf("Este es un mensaje de registro")
	log.Println("Este es otro mensaje de registro")
	log.Fatal("Soy un registro de errores") // detiene el
	log.Panic("Soy un panico")              // detiene el codigo y guarda el panico

	// Abrir archivos / crearlo / anadir contenido / solo escritura /  permisos
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("!Oye, soy un log")
}
