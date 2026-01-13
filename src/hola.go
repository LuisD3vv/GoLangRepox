package main

import (
	"fmt"
	"math"

	"rsc.io/quote"
)

// las constantes deben ser inicializadas por obvias razones
const Pi = 3.14

const (
	x = 100
	y = 0b1010 // binario
	z = 0o12   //octal
	w = 0xFF   // hexadecimal
)

const (
	/*iota nos sirve en declaraciones constantes de numeros incrementables, incrementandolo en uno o multiplicando*/
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

// otro ejemplo de iota.
type ByteSize float64

const (
	_           = iota // el primer valor se ignora asigandole un identificador vacio 0 cero
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("Hola mundo") // Escrito
	fmt.Println(quote.Go())   // Mediante el paquete

	// Declaracion de variables, no pueden quedarse in usar

	//var firstName, lastName string
	//var age int

	// mutivariables

	firstName, lastName, age := "Luis", "aguilar", 21
	/*:= nos sirve para inicializar nuevas variables sin usar var pero solo
	dentro de una funcion, declarando con var, podemos declararla ya sea
	dentro o fuera de la funcion*/

	// Al iniciarlizar la variable no es necesario colocarle el tipo de dato.

	fmt.Println(firstName, lastName, age)
	fmt.Println(Pi)

	fmt.Println(x, y, z, w)
	fmt.Println(Viernes)
	// enteros
	var integer int8 = -120
	var integer2 uint8 = 120
	fmt.Println(integer, integer2)
	// flotantes 64 es para numero muy grandes
	var float float32 = 12
	fmt.Println(float)
	// boleanos
	var valueBool bool = true // false
	if valueBool {
		fmt.Println("verdadero")
	}
	// string

	fullName := "Luis Alejandro \t(alias \"Lissandro\")\n"
	fmt.Println(fullName)
	// conocer el valor maximo y minimo que puede almacenar float e int
	fmt.Println(math.MaxUint8)

	// imprime el valor del codigo ascii / su valor decimal equivalente
	var a byte = 'a'
	fmt.Println(a)

	s := "hola"
	fmt.Println(s[0])

	//var r rune = '' valor unicode int de la cadena

}
