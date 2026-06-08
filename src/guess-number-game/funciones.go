package main

import "fmt"

func declaracion_funcion() {
	nombre := hello("Lissandro")
	fmt.Println(nombre)
	// con doble variable capturamos el doble resultado de la funcion
	suma, mult := calc(2, 3)
	fmt.Println("La suma es:", suma)
	fmt.Println("La multiplicacion es:", mult)
}

// debemos definir el tipo de dato de las variables y el de retorno
func hello(name string) string {
	return "Hola, " + name
}

// funcion con valores de retorno nombrados
func calc(a, b int) (sum, mult int) {
	sum = a + b
	mult = a * b
	return
}

// funcion con valores de retorno explicito
func calc2(a, b int) (int, int) {
	sum := a + b
	mult := a * b
	return sum, mult
}
