package main

import (
	"errors"
	"fmt"
)

func main() {
	horas, valorHora := 90.0, 5000.0
	salary, err := calcSalary(horas, valorHora)
	if err != nil {
		fmt.Println("\n", err)

	} else {
		if salary >= 150000 {
			salary -= salary * 0.1
		}
		fmt.Println("\nHoras Trabajadas: ", horas, "\nValor Hora: ", valorHora, "\nSalario: ", salary)
	}
}

func calcSalary(horas, valorHora float64) (result float64, err error) {
	if horas < 80 || horas < 0 {
		err = errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return 0, err
	}
	return horas * valorHora, nil
}
