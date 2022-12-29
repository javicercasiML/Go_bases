package main

import (
	"errors"
	"fmt"
)

type Cliente struct {
	Legajo    int
	Nombre    string
	Dni       string
	Telefono  string
	Domicilio string
}

var (
	ErrRepeat = errors.New("Error: el cliente ya existe")
)

var Clientes = []Cliente{ // o vacio var Products []Product
	{
		Legajo:    58004,
		Nombre:    "Javier Cercasi",
		Dni:       "41.999.999",
		Telefono:  "2619999999",
		Domicilio: "Avenida siempre viva",
	},
	{
		Legajo:    18564,
		Nombre:    "Juan Gomez",
		Dni:       "42.999.999",
		Telefono:  "2619999999",
		Domicilio: "Calle falsa 123",
	},
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Ejecución Finalizada.\n ")
	}()

	cliente := Cliente{1856, "Juan Gomez", "42.999.999", "2619999999", "Calle falsa 123"}
	cliente.existe(cliente.Legajo)
	cliente.validar(cliente)

}

func (c *Cliente) existe(leg int) {
	for _, cliente := range Clientes {
		if cliente.Legajo == leg {
			panic(ErrRepeat)
		}
	}
}

func (c *Cliente) validar(clt ...Cliente) {

	//print(ref)
	for _, dato := range clt {
		fmt.Println(dato, "rar")
	}
}
