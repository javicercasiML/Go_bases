package main

import (
	"errors"
	"fmt"
)

var (
	ErrSalary = errors.New("Error: el mínimo imponible es de 150.000")
)

func main() {
	var salary uint
	salary = 8000
	_, err := checkSalary(int(salary))

	if err != nil {
		fmt.Println(err, "\nResultado: ", errors.Is(err, ErrSalary))
	}
}

func checkSalary(s int) (result bool, err error) {
	if s <= 10000 {
		err = extra(s)
		return true, err
	}
	return result, nil
}

func extra(salary int) error {
	return fmt.Errorf("%w y el salario ingresado es de: %d", ErrSalary, salary)
}
