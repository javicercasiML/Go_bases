package main

import "fmt"

func main() {
	println("inicio")
	//aplicamos “defer” a la invocación de una función anónima
	defer func() {
		fmt.Println("Esta función se ejecuta a pesar de producirse panic")
	}()
	//creamos un panic con un mensaje de que se produjo
	println("Final")
	panic("se produjo panic!!!")
}
