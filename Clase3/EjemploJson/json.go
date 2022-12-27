package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	PrimerNombre string `json:"primer_nombre"`
	Apellido     string `json:"apellido"`
	Telefono     string `json:"telefono"`
	Direccion    string `json:"direccion"`
}

func main() {
	p := Persona{"Juan", "Perez", "43434343", "Calle falsa 123"}
	miJSON, err := json.Marshal(p)

	fmt.Println(string(miJSON))
	fmt.Println(err)
}
