package main

import (
	"errors"
	"fmt"
)

type Producto struct {
	Costo float64
}

type Pequeño struct {
	Producto
}

type Mediano struct {
	Producto
}

type Grande struct {
	Producto
	Adicional float64
}

type calculator interface {
	precio() float64
}

func (p *Pequeño) precio() float64 {
	return p.Costo
}

func (m *Mediano) precio() float64 {
	return m.Costo + m.Costo*0.03
}

func (g *Grande) precio() float64 {
	return g.Costo + g.Costo*0.06 + float64(g.Adicional)
}

const (
	pequeño = "Pequeño"
	mediano = "Mediano"
	grande  = "Grande"
)

func factoryProduct(tipo string, values ...float64) (instancia calculator, err error) {
	switch tipo {
	case pequeño:
		instancia = &Pequeño{Producto{Costo: values[0]}}
	case mediano:
		instancia = &Mediano{Producto{Costo: values[0]}}
	case grande:
		instancia = &Grande{
			Producto: Producto{Costo: values[0]}, Adicional: 2500}
	default:
		instancia = nil
		err = errors.New("\nProducto invalido")
	}
	return
}

func main() {

	productos := []string{pequeño, mediano, "Extra", grande}

	for _, producto := range productos {

		prod, err := factoryProduct(producto, 1000)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("\nEl precio del producto %s es de $%.2f \n", producto, prod.precio())
	}
	//p := factoryProduct(pequeño, 1000)
	//fmt.Printf("\nEl precio del producto %s es de $%.2f ", pequeño, p.precio())

	//m := factoryProduct(mediano, 1000)
	//fmt.Printf("\nEl precio del producto %s es de $%.2f", mediano, m.precio())

	//g := factoryProduct(grande, 1000, 2500)
	//fmt.Printf("\nEl precio del producto %s es de $%.2f\n\n", grande, g.precio())

}
