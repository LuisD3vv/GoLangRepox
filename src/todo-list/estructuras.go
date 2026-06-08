package main

import "fmt"

// estruct es un tipo de dato personalizado
type Persona struct {
	nombre string
	edad   int
	correo string
}

// se podria decir que es un metodo
// recibe un puntero de tipo persona necesita un receptor () y  el parametro dentro de ella
func (p *Persona) diHola() {
	// p es como de llamara a cada una de las "variables creadas a parti del objeto persona y con el punto podremos acceder a sus caracteristicas"
	fmt.Println("Hola mi nombre es", p.nombre)
}

func estructura() {
	Lissandro := Persona{
		nombre: "Luis",
		edad:   21,
		correo: "Lissandro@gmail.com"}

	fmt.Println(Lissandro.nombre)
	fmt.Println(Lissandro.edad)
	fmt.Println(Lissandro.correo)
	Lissandro.diHola()

	// punteros y metodos
	var x int = 10
	fmt.Println(x)
	editar(&x)
	fmt.Println(x)
}

/* recibe como argumento un puntero

Podemos percibirlo asi

& donde esta?
* que hay ahi?

por lo cual al modificar el valor mediante la direccion estanmos diciciendo


	var x int = 10
	sumar(&x) donde esta almacenado x?


	func sumar(x *int) {
	*x = x + 10 // lo que hay en esa direccion, ahora sumale 10
	}
*/

func editar(x *int) {
	*x = 20
}
