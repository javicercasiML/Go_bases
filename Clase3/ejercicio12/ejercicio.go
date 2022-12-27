package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (e *Employee) PrintEmployee() {
	fmt.Println("\nEmpleado de ID: ", e.ID, "\nNombre: ", e.Name, "\nPosicion: ", e.Position, "\nFecha de nacimiento: ", e.DateOfBirth)
}

func main() {
	empleado := Employee{
		Person:   Person{1, "Juan", "18-12-22"},
		ID:       1,
		Position: "Jefe",
	}
	empleado.PrintEmployee()

	empleado2 := Employee{
		Person:   Person{2, "Pedro", "18-12-98"},
		ID:       2,
		Position: "Secretario",
	}
	empleado2.PrintEmployee()

}
