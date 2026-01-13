///ackage main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

// recibe un puntero de tipo persona necesita un receptor () y  el parametro dentro de ella
func (p *Persona) diHola() {
	fmt.Println("Hola mi nombre es", p.nombre)
}

func main() {
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

// recibe como argumento un puntero
func editar(x *int) {
	*x = 20
}
