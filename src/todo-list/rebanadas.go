//package main

import "fmt"

// Matrices
func main() {
	// los 3 puntos es cuando no sabemos la cantidad para inicializar la matriz
	// basicamnet para que calcule automaticamente
	var a = [...]int{2, 4, 6, 8, 10}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	// forma con range
	for indice, valor := range a {
		fmt.Println(indice, valor)
	}

	// matriz bidimensional y su sintaxis.
	var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(matriz[i][j])
		}

	}
	fmt.Println(a)
	// slices
	diasSemana := []string{"Lunes", "Martes", "Miercoles",
		"Jueves", "Viernes", "Sabado", "Domingo"}
	fmt.Println(len(diasSemana))

	diaRebanada := diasSemana[0:5]
	fmt.Println(diaRebanada)

	diaRebanada = append(diaRebanada, "Sabado", "Domingo", "Otro dia")
	fmt.Println(diaRebanada)
	// los 3 puntos sirven para expandir un slice en valores individuales.
	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)
	fmt.Println(diaRebanada)

	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	//crear slices (rebanas)
	nombres := make([]string, 5, 10) // tipo, elementos y el maximo de capacidad
	nombres[0] = "Luis"
	fmt.Println(nombres)

	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)

	// copy regresa los elementos que se han copeado
	copy(rebanada2, rebanada1) // se aplica de derecha a izquierda
}
