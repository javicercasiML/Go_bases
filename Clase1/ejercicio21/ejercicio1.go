package main

import (
	"fmt"
)

func main() {
	palabra := "Argentina"
	Obtenerletras(palabra)
}

func Obtenerletras(palabra string) {
	fmt.Println("Cantidad de letras: ", len(palabra))
	for _, letra := range palabra {
		fmt.Println(string(letra))
	}

}
