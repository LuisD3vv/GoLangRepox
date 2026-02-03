package main

import (
	"fmt"
	"runtime"
	"time"
)

func queEsSwitch() {
	// podemos inicializar variables y usarlas dentro del if
	if t := time.Now(); t.Hour() > 12 {
		fmt.Println("manana")
	} else if t.Hour() > 17 {
		fmt.Println("Es tarde")
	} else {
		fmt.Println("Noche")
	}

	// condiciones dentro del switch
	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("manana")
	case t.Hour() < 17:
		fmt.Println("Es tarde")
	default:
		fmt.Println("Noche")
	}

	//switch, no es necsario colocarle el break y como vemos, pdoemos declarar la condicion ahi mismo
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("-> El sistema es windows")
	case "linux":
		fmt.Println("-> El sistema es linux")
	case "darwin":
		fmt.Println("-> El sistema es MacOS")
	default:
		fmt.Println("Otro Os")
	}
}
