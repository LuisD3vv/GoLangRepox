package main

import (
	"fmt"
	"os"
)

func UsoDeOS() {
	// difiere la ejecucion de una funcion

	// es este caso se agregan a un pila de ejecuicion LIFO
	// defer diferir el tiempo o llamda de una funcion

	// Ejemplo visual de ejecucion

	// A
	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// B
	defer file.Close()

	//C
	_, err = file.Write([]byte("Hola, Lissandro"))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	//salida

	//defer dice, pase lo que pase cierro el archivo

	/*
		order con defer:
			A - crear
			C - escribir
			B - pase lo que pase se cierra
	*/

}
