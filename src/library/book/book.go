package book

import "fmt"

// las funciones y metodos son asi
// Mayusculas - > Publicos(exportados) Minusculas -> privados al paquete

type Printable interface {
	// define conjunto de metodos, pero no su implementacion
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	title  string
	author string
	pages  int
}

// Contructor

func NewBook(title, author string, pages int) *Book {
	// Devuleve la direccion del libro, no el libro
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

// desferenciar = ir a la direccion y usar lo que hay ahi

/*
Puntero = donde esta
desreferencia = que hay ahi
*/

// Simulando encapsulamiento

// utilizando get y set para atributos privados del objeto original por eso el puntero
func (b *Book) SetTitle(title string) {
	b.title = title
}
func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAutor: %s\nPages: %d\n",
		b.title, b.author, b.pages)
}

type TextBook struct {
	Book
	editorial string
	levels    string
}

func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		levels:    level,
	}
}
func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\nAutor: %s\nPages: %d\nEditorial: %s\nLevel: %s\n",
		b.title, b.author, b.pages, b.editorial, b.levels)
}
