package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDogFood es un test unitario que comprueba
// el correcto funcionamiento de la funcion dogFood
func TestDogFood(t *testing.T) {
	//arange
	var amount float64
	animal, cantidad, expectedRes := dog, 10, 200.0

	//act
	cant, err := Animal(animal)
	amount += cant(cantidad)
	amount += cant(cantidad)

	//assert
	assert.Equal(t, expectedRes, amount)
	assert.Equal(t, nil, err)
}

// TestCatFood es un test unitario que comprueba
// el correcto funcionamiento de la funcion catFood
func TestCatFood(t *testing.T) {
	//arange
	var amount float64
	animal, cantidad, expectedRes := cat, 10, 100.0

	//act
	cant, err := Animal(animal)
	amount += cant(cantidad)
	amount += cant(cantidad)

	//assert
	assert.Equal(t, expectedRes, amount)
	assert.Equal(t, nil, err)
}

// TestHamsterFood es un test unitario que comprueba
// el correcto funcionamiento de la funcion hamsterFood
func TestHamsterFood(t *testing.T) {
	//arange
	var amount float64
	animal, cantidad, expectedRes := hamster, 10, 5.0

	//act
	cant, err := Animal(animal)
	amount += cant(cantidad)
	amount += cant(cantidad)

	//assert
	assert.Equal(t, expectedRes, amount)
	assert.Equal(t, nil, err)
}

// TestSpiderFood es un test unitario que comprueba
// el correcto funcionamiento de la funcion spiderFood
func TestSpiderFood(t *testing.T) {
	//arange
	var amount float64
	animal, cantidad, expectedRes := spider, 10, 3.0

	//act
	cant, err := Animal(animal)
	amount += cant(cantidad)
	amount += cant(cantidad)

	//assert
	assert.Equal(t, expectedRes, amount)
	assert.Equal(t, nil, err)
}

// TestMixFood es un test unitario que comprueba
// el correcto funcionamiento de la funcion spiderFood y dogFood
func TestMixFood(t *testing.T) {
	//arange
	var amount float64
	animal, animal2, cantidad, cantidad2, expectedRes := spider, dog, 50, 10, 107.5

	//act
	cant, _ := Animal(animal)
	cant2, _ := Animal(animal2)
	amount += cant(cantidad)
	amount += cant2(cantidad2)

	//assert
	assert.Equal(t, expectedRes, amount)
}

// TestAnimalError es un test unitario que comprueba
// el funcionamiento de la funcion Animal, con un tipo de
// animal inexistente
func TestAnimalError(t *testing.T) {
	//arange
	animal := "pato"
	expectedErr := errors.New("No existe el animal")

	//act
	cant, err := Animal(animal)

	//assert
	assert.Nil(t, cant)
	assert.Equal(t, expectedErr, err)
}
