//package main

import "fmt"

func main() {
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

func calc(a, b int) (sum, mult int) {
	sum = a + b
	mult = a * b
	return
}
