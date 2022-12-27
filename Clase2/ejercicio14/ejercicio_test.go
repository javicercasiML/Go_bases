package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMinimum es un test unitario que comprueba
// el correcto funcionamiento de la funcion minimunFunction
// (encuentra el minimo valor de un conjunto).
func TestMinimum(t *testing.T) {
	//arange
	tipo := minimum
	notas, expectedRes := []float64{2, 3, 3, 4, 10, 2, 4, 5}, 2.0

	//act
	prom, err := operation(tipo)
	result := prom(notas...)

	//assert
	assert.Equal(t, expectedRes, result)
	assert.Equal(t, nil, err)
}

// TestMaximun es un test unitario que comprueba
// el correcto funcionamiento de la funcion maximunFunction
// (encuentra el valor mas alto de un conjunto).

func TestMaximun(t *testing.T) {
	//arange
	tipo := maximum
	notas, expectedRes := []float64{2, 3, 3, 4, 10, 2, 4, 5}, 10.0

	//act
	prom, err := operation(tipo)
	result := prom(notas...)

	//assert
	assert.Equal(t, expectedRes, result)
	assert.Equal(t, nil, err)
}

// TestAverage es un test unitario que comprueba
// el correcto funcionamiento de la funcion averageFunction
// (encuentra el promedio de un conjunto de valores).
func TestAverage(t *testing.T) {
	//arange
	tipo := average
	notas, expectedRes := []float64{2, 3, 3, 4, 10, 2, 4, 5}, 4.125

	//act
	prom, err := operation(tipo)
	result := prom(notas...)

	//assert
	assert.Equal(t, expectedRes, result)
	assert.Equal(t, nil, err)
}

// TestError es un test unitario que comprueba
// el funcionamiento de la funcion operation, con una
// categoria inexistente
func TestError(t *testing.T) {
	//arange
	tipo := "prom"
	expectedErr := errors.New("Operacion invalida")
	//act
	prom, err := operation(tipo)

	//assert
	assert.Nil(t, prom)
	assert.Equal(t, expectedErr, err)
}
