package main

import "library/animal"

func main() {
	// Creando el objeto
	// mytextBook := book.NewTextBook("ingles", "Eduardo Aguilar", 261, "Santillana SAC", "Secundaria")
	//mytextBook.PrintInfo()
	// book.Print(mytextBook)

	animales := []animal.Animal{
		&animal.Gato{Nombre: "Eliot"},
		&animal.Gato{Nombre: "Benito"},
		&animal.Gato{Nombre: "Mike"},
		&animal.Perro{Nombre: "Rocko"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
