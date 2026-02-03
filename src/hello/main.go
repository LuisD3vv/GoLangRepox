package main

import (
	"fmt"
	"log"

	"github.com/LuisD3vv/greetings"
)

func main() {
	log.SetPrefix("greetings: ") // prefijo que tendra cada mensaje de log
	log.SetFlags(0) // no incluye fecha y hora. solo el mensaje

	nombres := []string{"Luis", "Eduardo", "Wiliam"}
	messages, err := greetings.Hellos(nombres)
	if err != nil {
		log.Fatal(err)
	}

	message, err := greetings.Hello("Lissandro")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	fmt.Println(messages)
}

// go mod tidy instala faltantes y elimina no usados
