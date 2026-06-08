package main

import "fmt"

func loops() {
	// bucle declarado
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	// bucle infinito
	for {
		fmt.Println("Hola,soy infinito")
		break
	}

	// bucle infinito con break y condicion
	n := 10
	for n < 0 {
		if n == 0 {
			break
		}
		fmt.Println("Hola, estoy decrementando", n)
		n -= 1
	}
	// con condicion boleana
	n2 := 0
	for n2 < 0 {
		if n2 == 10 {
			break
		}
		fmt.Println("Hola, estoy decrementando", n2)
		n2 += 1
	}

	// for con un puto iterable
	numeros := [...]int{1, 2, 3, 4, 5}

	for numero := range numeros {
		fmt.Println("Color: ", numero)
	}
	// con strings
	vocales := [...]string{"a", "e", "i", "o", "u"}
	for vocal := range vocales {
		fmt.Println("Color: ", vocal)
	}
}
