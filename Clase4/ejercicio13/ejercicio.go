package main

import (
	"errors"
	"fmt"
)

var (
	ErrSalary = errors.New("Error: el salario es menor a 10.000")
)

func main() {
	var salary uint
	salary = 8000
	_, err := checkSalary(int(salary))

	if err != nil {
		fmt.Println("\nResultado: ", errors.Is(err, ErrSalary))
	}
}

func checkSalary(s int) (result bool, err error) {
	if s <= 10000 {
		err = ErrSalary
		return true, err
	}
	return result, nil
}
