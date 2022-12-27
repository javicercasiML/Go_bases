package main

import (
	"fmt"
)

func main() {
	prestamo(23, true, 5, 120000)
	prestamo(20, true, 5, 120000)
	prestamo(23, true, 5, 90000)
}

func prestamo(edad int, empleado bool, antig int, sueldo int) {

	interes := true
	if edad > 22 && empleado && antig > 1 {
		if sueldo > 100000 {
			interes = false
		}
		fmt.Println("\nEmpleado: ", empleado, " Edad: ", edad, " Antiguedad: ", antig, " Prestamo: Aceptado", "Interes: ", interes)
	} else {
		fmt.Println("\nEmpleado: ", empleado, " Edad: ", edad, " Antiguedad: ", antig, " Prestamo: Rechazado")
	}
}
