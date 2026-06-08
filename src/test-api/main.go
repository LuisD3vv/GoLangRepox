package main

// Concurrencia anal

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	// Lista de servidores de apiss
	apis := []string{
		"https://managment.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string) // Los canales propocionan una forma de sincronizar y comunicar datos entre goroutines
	// valor := <-ch           // recibir valor de un canal a una sola variable o puede ser un slice

	// devuelve indice y valor
	for _, api := range apis {
		go checkAPI(api, ch) // goroutines / concurrencia, es un hilo de ejecucion gestionado por el entorno de ejecucion.
	}

	// recoge los resultado del canal, por eso el len(), con len ya sabemos cuantos resultados esperar
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	fmt.Println(<-ch)

	elapsed := time.Since(start)
	fmt.Printf("Listo, tomo %v segundos!\n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("ERROR : %s esta caido!\n", err)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: !%s esta en funcionamiento\n", api)

}
