package main

import "fmt"

// mapas

func queSonMaps() {
	// elementos clave valor, por defecto se acomodan el orden alfabetico
	colors := map[string]string{ // tipo de dato de la clave y el valor
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
	fmt.Println(colors["rojo"])
	colors["negro"] = "#000000"
	fmt.Println(colors)

	// ok es una comprobacion que dependiendo si clave existe devuelve true y false en viceversa.

	if valor, ok := colors["rojo"]; ok {
		fmt.Println(valor)
	} else {
		fmt.Println("no existe la clave")
	}

	delete(colors, "rojo")
	fmt.Println(colors)

	for clave, valor := range colors {
		fmt.Printf("clave: %s, Valor: %s\n", clave, valor)
	}
}
