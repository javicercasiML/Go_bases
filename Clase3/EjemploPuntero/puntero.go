package main

import "fmt"

// Increase recibe un puntero de tipo entero
func Increase(v *int) {
	// Desreferenciamos la variable v para obtener
	// su valor e incrementarlo en 1
	*v++
}
func main() {
	var v int = 19
	// La función Increase recibe un puntero
	// utilizamos el operador de dirección &
	// para pasar la dirección de memoria de v
	Increase(&v)
	fmt.Println("El valor de v ahora vale:", v)
}
