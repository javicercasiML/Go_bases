package main

import (
	"errors"
	"fmt"
)

var (
	ErrSalary = &Myerror{"El salario es menor a 10.000"}
)

type Myerror struct {
	text string
}

// Que el type struct implemente el metodo Error()
func (e *Myerror) Error() string {
	return fmt.Sprintf("Error: %s", e.text)
}

func main() {
	var salary uint
	salary = 8000
	_, err := checkSalary(int(salary))

	if errors.Is(err, ErrSalary) {
		fmt.Println(err, "\nResultado del error: ", errors.Is(err, ErrSalary))
	} else {
		fmt.Println("El salario es: $", salary)
	}

}

func checkSalary(s int) (result bool, err error) {
	if s <= 10000 {
		return true, ErrSalary
	}
	return result, nil
}
