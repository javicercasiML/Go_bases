package main

import "fmt"

func isPair(num int) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}

	}()

	if (num % 2) != 0 {
		panic("no es un número par")
	}

	fmt.Println(num, " es un número par!")
}

func main() {
	isPair(4)
	fmt.Println("Ejecución completada!")
}
