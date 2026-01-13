package main

import (
	"fmt"
	"math"
)

func main() {

	var lado_1 float64
	var lado_2 float64
	var hipotenusa float64
	var area float64
	var perimetro float64
	fmt.Println("--= Calcular el area y el perimetro de un triangulo =--")

	fmt.Print("Ingrese el lado 1: ")
	fmt.Scanln(&lado_1)

	fmt.Print("Ingrese el lado 2: ")
	fmt.Scanln(&lado_2)

	hipotenusa = math.Sqrt(math.Pow(lado_1, 2) + math.Pow(lado_2, 2))
	perimetro = lado_1 + lado_2
	area = lado_1 * lado_2 / 2
	/*
		La base puede ser cualquiera de los dos catetos
		no hay uno oficialmente base y otro altura
		La base y altura siempre son perpendiculares entre si
	*/
	fmt.Printf("La hipotenusa es: %.2f\n", hipotenusa)
	fmt.Printf("El perimetro es: %.2f\n", perimetro)
	fmt.Printf("El area es: %.2f\n", area)
}
