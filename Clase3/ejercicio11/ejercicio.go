package main

import "fmt"

type Product struct {
	ID          int
	Nombre      string
	Price       string
	Description string
	Category    string
}

var Products = []Product{ // o vacio var Products []Product
	{
		ID:          0,
		Nombre:      "Monitor",
		Price:       "10000",
		Description: "Monitor Led Samsung",
		Category:    "Computacion",
	},
	{
		ID:          4,
		Nombre:      "Monitor",
		Price:       "10000",
		Description: "Monitor Led Noblex",
		Category:    "Computacion",
	},
}

func main() {
	producto := Product{3, "Lavarropas", "200000", "Samsung 8Kg", "Hogar"}
	producto.Save(producto)
	fmt.Println(Products)
	producto.GetAll()
	getById(4)
}

func (p *Product) Save(prods ...Product) {
	for _, prod := range prods {
		Products = append(Products, prod)
	}
	//
}

func (p *Product) GetAll() {
	for i, prodd := range Products {
		fmt.Println("El producto ", i, " es: ", prodd)
	}
	//
}

func getById(id int) {

	for _, prod := range Products {
		if id == prod.ID {
			fmt.Println("El producto buscado de ID ", id, " es: ", prod)
		}
	}
}
