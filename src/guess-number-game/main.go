package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()
}
func play() {
	num_aletorio := rand.Intn(100)
	var numero_usuario int
	var intentos int
	const maxIntentos = 10

	fmt.Println("Guess the number")
	fmt.Println("Intenta adivinar el numero secreto.")
	for intentos < maxIntentos {
		fmt.Printf("\nIngresa el numero (intentos restantes: %d)\n", maxIntentos-intentos)
		fmt.Scanln(&numero_usuario)

		if numero_usuario == num_aletorio {
			fmt.Printf("Felicidades!!, el numero era %d: ", num_aletorio)
			jugarNuevamente()
			break
		}

		if numero_usuario > num_aletorio {
			fmt.Println("el Numero secreto es menor")
			intentos++
		} else if numero_usuario < num_aletorio {
			fmt.Println("el Numero secreto es mayor")
			intentos++
		}
	}
	fmt.Printf("Te has quedado sin intentos!, el numero era: %d\n", num_aletorio)
	jugarNuevamente()
}
func jugarNuevamente() {
	var eleccion string
	fmt.Println("Quieres jugar Nuevamente? (s/n)")
	fmt.Scanln(&eleccion)
	switch eleccion {
	case "s":
		play()
	case "n":
		fmt.Println("Gracias por jugar")
	default:
		fmt.Println("Elecciojn invalida")
		jugarNuevamente()
	}
}
