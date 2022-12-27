package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCategoriaA es un test unitario que comprueba
// el correcto funcionamiento de la funcion salario, categoria A
func TestCategoriaA(t *testing.T) {
	//arange
	var minutos, expectedRes float64 = 120, 2000
	categoria := "C"

	//act
	prom, err := salario(minutos, categoria)

	//assert
	assert.Equal(t, expectedRes, prom)
	assert.Equal(t, nil, err)
}

// TestCategoriaB es un test unitario que comprueba
// el correcto funcionamiento de la funcion salario, categoria B
func TestCategoriaB(t *testing.T) {
	//arange
	var minutos, expectedRes float64 = 180, 5400
	categoria := "B"

	//act
	prom, err := salario(minutos, categoria)

	//assert
	assert.Equal(t, expectedRes, prom)
	assert.Equal(t, nil, err)
}

// TestCategoriaC es un test unitario que comprueba
// el correcto funcionamiento de la funcion salario, categoria C
func TestCategoriaC(t *testing.T) {
	//arange
	var minutos, expectedRes float64 = 300, 5000
	categoria := "C"

	//act
	prom, err := salario(minutos, categoria)

	//assert
	assert.Equal(t, expectedRes, prom)
	assert.Equal(t, nil, err)
}

// TestCategoriaD es un test unitario que comprueba
// el funcionamiento de la funcion salario, con una
// categoria inexistente
func TestCategoriaD(t *testing.T) {
	//arange
	var minutos, expectedRes float64 = 300, 0
	categoria := "D"
	expectedErr := errors.New("Categoria inexistente")

	//act
	prom, err := salario(minutos, categoria)

	//assert
	assert.Equal(t, expectedRes, prom)
	assert.Equal(t, expectedErr, err)
}
