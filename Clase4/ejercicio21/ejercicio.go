package main

import (
	"fmt"
	"os"
)

func main() {
	lectura("customers.txt")
	fmt.Println("Ejecución Finalizada.\n ")
}

func lectura(path string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	file, err := os.ReadFile(path)
	data := string(file)
	fmt.Println("\nDatos del archivos:", data, "\n ")

	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

}
