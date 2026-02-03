package main

import (
	"fmt"
	"strconv"
)

func conversion() {
	var integer16 int16 = 50
	var integer32 int32 = 100
	// casteo explicito

	s := "100"
	// como devuelve dos valores, podemos seleccionar cual queremos utilizar
	i, _ := strconv.Atoi(s) // strings a numero (si es posible) ejemplo "1" o "2" no "a" o puede conversion tabla ascii
	fmt.Println(i + i)
	n := 42
	s = strconv.Itoa(n) // numero a strings
	fmt.Println(s)
	// no se pueden realizar operaciones entre tipos,es obvio
	fmt.Println(integer16 + int16(integer32))

	name := "Luis"
	age := 21
	// tal y como funciona en el lenguaje C
	fmt.Printf("hola, me llamo %s y tengo %d annos\n", name, age)

	greeting := fmt.Sprintf("hola, me llamo %s y tengo %d annos", name, age)
	fmt.Println(greeting)

	// ingresar usuarios

	var nombre string
	var edad int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&nombre)

	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&edad)

	fmt.Printf("hola, me llamo %s y tengo %d annos\n", nombre, edad)

	// conocer el tipo de dato

	fmt.Printf("El tipo de nombre es %2.T", nombre)
	fmt.Printf("El tipo de edad es %2.T", edad)
}
