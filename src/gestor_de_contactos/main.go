package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divideManejoErrores(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		// con este paquete podemos crear errores personalizados, ya que en go los errores se manejan con if
		return 0, errors.New("No se puede dividir por cero.")

	}
	return dividendo / divisor, nil
}

// Manejo de errores
func Manejo_de_errores() error {
	resultado, err := divideManejoErrores(3, 2)
	if err != nil {
		return err

	}
	fmt.Println("Resultado:", resultado)
	str := "123"
	numero, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	fmt.Println("Numero: ", numero)

	return nil
}
