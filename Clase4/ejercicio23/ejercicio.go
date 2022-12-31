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
	ErrRepeat = errors.New("El cliente ya existe")
	ErrEmpty  = errors.New("No se permiten registros vacios")
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
			fmt.Println("\nSe detectaron varios errores en tiempo de ejecución:", err)
		}
		fmt.Println("\nFin de la ejecución\n ")
	}()

	cliente := Cliente{1856, "Juan Gomez", "42.999.999", "261999999", "Calle falsa 123"}
	_, err := cliente.existe(cliente.Legajo)
	if err != nil {
		panic(err)
	}
	err = cliente.validar(cliente)
	if err != nil {
		panic(err)
	}

	Clientes = append(Clientes, cliente)

}

func (c *Cliente) existe(leg int) (val int, err error) {
	for _, cliente := range Clientes {
		if cliente.Legajo == leg {
			return 0, ErrRepeat
		}
	}
	return
}

func (c *Cliente) validar(clt ...Cliente) (err error) {

	for _, c := range clt {
		if c.Legajo != 0 && c.Dni != "" && c.Domicilio != "" && c.Nombre != "" && c.Telefono != "" {
			return
		}
	}
	return ErrEmpty
}
