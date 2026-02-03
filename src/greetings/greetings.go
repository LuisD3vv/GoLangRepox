package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello devuelve un saludo

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre Vacio")
	}
	// devuleve un saludo a la persona
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hello regresa un mapa que asciosa cada nombre del slice con un saludo
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := [...]string{
		"Hola, %v !Bienvenido.",
		"!Que bueno verte %v.",
		"!Me da gusto verte %v.",
		"Listo? %v.",
		"Bonjour %v!.",
		"Encantado de conocerte %v.",
	}
	return formats[rand.Intn(len(formats))]
}
