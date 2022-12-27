package main

import (
	"errors"
	"fmt"
)

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	spider  = "spider"
)

func dogFood(num int) (result float64) {
	return float64(num) * 10
}

func catFood(num int) (result float64) {
	return float64(num) * 5
}

func hamsterFood(num int) (result float64) {
	return float64(num) * 0.25
}

func spiderFood(num int) (result float64) {
	return float64(num) * 0.15
}

type AnimalFunction func(num int) float64

func Animal(anml string) (fn AnimalFunction, err error) {
	switch anml {
	case dog:
		return dogFood, err
	case cat:
		return catFood, err
	case hamster:
		return hamsterFood, err
	case spider:
		return spiderFood, err
	default:
		err = errors.New("No existe el animal")
		return nil, err
	}
}
func main() {

	animales := []string{dog, cat, hamster, spider, "bunny"}
	cantidades := []int{10, 20, 5, 5, 10}
	var amount float64

	for i, animal := range animales {
		fun, err := Animal(animal)

		if err != nil {
			fmt.Println(err, animal)
		} else {
			amount += fun(cantidades[i])
		}
	}
	fmt.Print("La cantidad de alimento necesaria es: ", amount, "Kg.\n\n")

	//animalDog, msg := animal(dog)
	//animalCat, msg := animal(cat)
	//
	//var amount float64
	//amount += animalDog(10)
	//amount += animalCat(10)

}
