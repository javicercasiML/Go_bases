package main

import (
	"fmt"
)

func main() {
	mes(1)
	mes(3)
	mes(15)
}

func mes(num int) {

	meses := map[int]string{1: "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre"}

	if num > 0 && num <= 12 {
		fmt.Println("El mes solicitado es: ", meses[num])
	} else {
		fmt.Println("Mes invalido")
	}
}
