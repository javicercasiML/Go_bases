package main

import (
	"errors"
	"fmt"
)

func main() {
	prom, err := promedio(1, 5, -9, 10)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("El promedio es: ", prom)
	}
}

func promedio(notas ...float64) (prom float64, err error) {
	var suma float64
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("No se permiten numeros negativos")
		}
		suma += nota
	}
	return suma / float64(len(notas)), err
}
