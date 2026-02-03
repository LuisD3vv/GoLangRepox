package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Struct tipo producto, aqui lo guardamos en un archivo
type Producto struct {
	Nombre string  `json:"nombre"`
	Precio float32 `json:"precio"`
	Stock  int     `json:"stock"`
}

func main() {
	menu()
}

func SaveProducts(Productos []Producto) error {
	file, err := os.Create("Stock.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(Productos)
	if err != nil {
		return err
	}
	return nil
}

func LoadProducts(Productos *[]Producto) error {
	data, err := os.ReadFile("Stock.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, Productos)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func DeleteProducts(productos []Producto) error {
	// hay que tener cuidado al crear slices asi, ya que la longitud debe estar en 0 para no teener elementos basura, pero la capacidad si la podemos predefinir
	temslice := make([]Producto, 0)
	// hay que tener cuidado, porque podemos crear indices vacios o con basura, al crear el largo o capacidad con un valor no verificado previamente
	var deleteIndex int
	fmt.Println("=========== Productos disponibles ==========")
	for index, producto := range productos {
		fmt.Printf("%d. Nombre: %s precio: $%.2f Stock: %d\n",
			index+1, producto.Nombre, producto.Precio, producto.Stock)
	}

	fmt.Println("Ingresa el indice del producto que deseas eliminar")
	// antes me daba error porque no especifioque el tipo de dato que leeria jeje
	_, err := fmt.Scanf("%d", &deleteIndex)
	if err != nil {
		fmt.Println("Porfavor, ingresa un indice valido: ", err)
		return err
	}

	fmt.Println(deleteIndex)
	// utilizar indice real es decir -1
	for index, producto := range productos {
		if index == deleteIndex-1 {
			fmt.Println("Producto eliminado: ", producto.Nombre)
			continue
		}
		temslice = append(temslice, producto)
	}

	data, err := json.MarshalIndent(temslice, "", " ")
	/*
		json.mashall function from the standard encoding/json package is used to convert Go data structures (like structs, maps, and slices) into a JSON-encoded byte slice.
	*/
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("Stock.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func AddProducts(productos []Producto) {
	reader := bufio.NewReader(os.Stdin)
	var p Producto
	fmt.Print("Nombre del producto: ")
	p.Nombre, _ = reader.ReadString('\n') // al guardarse al final se le pone un salto, al igual que imprimirlo de vuelta.
	fmt.Print("Precio: ")
	fmt.Scanln(&p.Precio)
	fmt.Print("Stock: ")
	fmt.Scanln(&p.Stock)
	// Agregar un producto al slice, mediante su metodo .append(slice,item)
	productos = append(productos, p)

	fmt.Println("Producto Agregado Correctamente")
	if err := SaveProducts(productos); err != nil {
		fmt.Println("Hubo un error al guardar el producto")
	}
}

func ModProduct(productos []Producto) {
	reader := bufio.NewReader(os.Stdin)
	var p Producto
	var Modindex int
	ListProducts(productos)
	fmt.Println("Menu de edicion de producto\nA continuacion ingrese el indice el producto")
	fmt.Scanln(&Modindex)
	// Mostrar el producto a modificar
	fmt.Print("Nombre del producto: ")
	p.Nombre, _ = reader.ReadString('\n') // al guardarse al final se le pone un salto, al igual que imprimirlo de vuelta.
	fmt.Print("Precio: ")
	fmt.Scanln(&p.Precio)
	fmt.Print("Stock: ")
	fmt.Scanln(&p.Stock)

}

func Sale(productos []Producto) {

	descuento := 0.15
	var total_sin_descuento float32
	var total_con_descuento float32
	var total float32
	var cantidad int32
	var indice_producto int32

	// slice de productos
	// los slices son estructuras de datos dinamicas, muy utiles para guardar o eliminar datos en ejecucion
	for {
		fmt.Println("=========== Productos disponibles ==========")
		for index, producto := range productos {
			fmt.Printf("%d. Nombre: %s precio: $%.2f Stock: %d\n", index, producto.Nombre, producto.Precio, producto.Stock)
		}
		fmt.Println("============================================")
		fmt.Printf("Subtotal: $%.2f\n", total_sin_descuento)
		fmt.Println("Escribe el indice del producto que deseas llevar (-1 para salir)")
		fmt.Scanln(&indice_producto)

		if indice_producto < 0 {
			if total_sin_descuento == 0 {
				fmt.Println("No has comprado nada weon")
				return
			} else if total_sin_descuento >= 750 {
				total_con_descuento = total_sin_descuento * float32(descuento)
				total = total_sin_descuento - total_con_descuento
				fmt.Printf("Felicidades Aplicaste para un descuento del 15 porciento ahora pagas: $%.2f\n", total)
				fmt.Printf("Con un descuento del: $%.2f\n", total_con_descuento)
				comprarDeNuevo()
				return
			} else {
				fmt.Printf("No aplicaste para un descuento, el total es: $%.2f\n", total_sin_descuento)
				comprarDeNuevo()
			}
		}
		if indice_producto > 0 && indice_producto < int32(len(productos)) {
			total_sin_descuento += productos[indice_producto].Precio * float32(cantidad)
		} else {
			fmt.Printf("Porfavor ingresa un numero dentro del menu.\n")
			continue
		}
		fmt.Printf("%s", productos[indice_producto].Nombre)
		fmt.Println("Cuantos vas a llevar?: ")
		fmt.Scanln(&cantidad)
	}
}

func ListProducts(productos []Producto) {
	fmt.Println("============ Todos los Productos ===========")
	for index, producto := range productos {
		fmt.Printf("%d. Nombre: %s precio: $%.2f Stock-inStore: %d\n", index+1, producto.Nombre, producto.Precio, producto.Stock)
	}
	fmt.Println("============================================")
}

func menu() {
	var productos []Producto //por valor es decir una copia al pasarse
	//var productos2 *[] por referencia  de memoria
	// var productos []Producto
	LoadProducts(&productos) // cargar productos

	opciones := [...]string{"Vender", "Agregar Producto", "Eliminar Producto", "Listar Producto", "Modificar Productos", "Salir"}

	for {
		fmt.Println("=====-- SOFT TIENDITA V7 --=====")
		for index, opcion := range opciones {
			fmt.Printf("%d- %s\n", index+1, opcion)
		}
		var menuOption int
		_, err := fmt.Scanln(&menuOption)
		if err != nil {
			fmt.Println("error al leer la opcion")
		}
		switch menuOption {
		case 1:
			Sale(productos)
		case 2:
			AddProducts(productos)
		case 3:
			DeleteProducts(productos)
		case 4:
			ListProducts(productos)
		case 5:
			ModProduct(productos)
		case 6:
			fmt.Println("Hasta pronto!")
			return
		default:
			fmt.Println("Opcion no valida")
		}
	}
}

func comprarDeNuevo() {
	var eleccion string
	fmt.Print("Quieres realizar algo Mas? (s/n): ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		menu()
	case "n":
		fmt.Println("Gracias por comprar en la tiendita")
		return
	default:
		comprarDeNuevo()
	}
}
