package animal

import "fmt"

type Animal interface {
	Sonido()
}

// Estructura de perro y sus metodos
type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " Guau guau papu")
}

// Estructura de perro y sus metodos
type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + " Miau Miau papu")
}

// Funcion para imprimir el sonido

func HacerSonido(animal Animal) {
	animal.Sonido()
}
