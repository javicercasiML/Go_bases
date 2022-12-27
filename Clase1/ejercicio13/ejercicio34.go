package main

import (
	"fmt"
)

// Ejercicio 3:
var nombre string
var apellido string
var edad int
var licencia = true
var estatura_de_la_persona int

// Ejercicio 4:
var apell string = "Gomez"
var edad1 int = 35
var sueldo float64 = 45857.90
var nombre1 string = "Julian"

func main() {
	apellido1 := "apellido" // Para funciones locales
	boolean := false        // Para funciones locales
	fmt.Println(boolean, apellido1)
	fmt.Println(nombre, apellido, edad, licencia, estatura_de_la_persona, apell, edad1, sueldo, nombre1)
}
