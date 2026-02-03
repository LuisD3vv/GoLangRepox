package main

import "fmt"

func divide(dividendo, divisor int) {
	// funcion anonima que captura el panico, de esta forma no interfiere con el codigo
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}() // para que se llame sola
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No puedes dividir por cero") // provocar un error
	}
}

func introducir_datos() {
	divide(100, 10)
	divide(100, 0)
	divide(200, 25)
}
