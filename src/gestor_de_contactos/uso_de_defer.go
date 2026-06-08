package main

import (
	"os"
)

func UsoDeOS() error {
	// difiere la ejecucion de una funcion

	// es este caso se agregan a un pila de ejecuicion LIFO
	// defer diferir el tiempo o llamda de una funcion

	// Ejemplo visual de ejecucion

	// A
	file, err := os.Create("hola.txt")
	if err != nil {
		return err
	}

	// C
	defer file.Close()

	// B
	_, err = file.Write([]byte("Hola, Lissandro"))
	if err != nil {
		return err
	}

	//salida

	//defer dice, pase lo que pase cierro el archivo o hago lo que dice

	/*
		order con defer:
			A - crear
			C - pase lo que pase se cierra
			B - escribir
	*/
	return nil
}
