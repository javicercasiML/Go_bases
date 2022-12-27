package main

import "fmt"

func main() {
	fmt.Println("\nEl empleado tendra un impuesto de: ", impuesto(40000))
	fmt.Println("\nEl empleado tendra un impuesto de: ", impuesto(50000))
	fmt.Println("\nEl empleado tendra un impuesto de: ", impuesto(160000))
}

func impuesto(sueldo uint) string {
	if sueldo < 50000 {
		return "0%"
	} else if sueldo >= 50000 && sueldo < 150000 {
		return "17%"
	}
	return "27%"
}
