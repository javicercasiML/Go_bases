package main

import "fmt"

type Myerror struct {
	text string
}

func main() {
	var salary uint
	salary = 200000
	if salary < 150000 {
		err := &Myerror{"El salario ingresado no alcanza el minimo imponible"}
		fmt.Println(err)
	} else {
		err := &Myerror{"Debe pagar impuestos"}
		fmt.Println(err)
	}
}

// Que el type struct implemente el metodo Error()
func (e *Myerror) Error() string {
	return fmt.Sprintf("Error: %s", e.text)
}
