package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a *Alumnos) Detalle() {
	fmt.Println("\nNombre: ", a.Nombre, "\nApellido: ", a.Apellido, "\nDNI: ", a.DNI, "\nFecha: ", a.Fecha)
}

func main() {
	alumno := Alumnos{"Javier", "Cercasi", "41200300", "26-12-22"}
	alumno.Detalle()

}
