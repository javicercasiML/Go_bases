package main

import (
	"errors"
	"fmt"
)

var categorias = map[string]float64{"C": 1000, "B": 1500, "A": 3000}

func main() {
	sal, _ := salario(120, "C")
	fmt.Println("Su salario es: ", sal)
	sal, _ = salario(180, "B")
	fmt.Println("Su salario es: ", sal)
	sal, _ = salario(90, "B")
	fmt.Println("Su salario es: ", sal)
}

func salario(minutos float64, categoria string) (resultado float64, err error) {
	horas := minutos / 60
	resultado = horas * categorias[categoria]
	switch categoria {
	case "A":
		resultado += resultado * 0.5
	case "B":
		resultado += resultado * 0.2
	case "C":
		resultado = resultado * 1
	default:
		resultado = 0
		err = errors.New("Categoria inexistente")
	}
	return
}
