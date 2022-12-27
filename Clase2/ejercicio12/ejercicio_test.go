package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromedioCorrecto(t *testing.T) {
	//arange
	notas := []float64{1, 5, 9, 10}
	expectedProm := 6.25

	//act
	prom, err := promedio(notas...)

	//assert
	assert.Equal(t, expectedProm, prom)
	assert.Equal(t, nil, err)
}

func TestPromedioCorrect(t *testing.T) {
	//arange
	notas := []float64{10, 10, 9, 10}
	expectedProm := 9.75

	//act
	prom, err := promedio(notas...)

	//assert
	assert.Equal(t, expectedProm, prom)
	assert.Equal(t, nil, err)
}

func TestPromedioIncorrecto(t *testing.T) {
	//arange
	notas := []float64{10, -10, 9, 10}
	expectedProm := 0.0
	expectedErr := errors.New("No se permiten numeros negativos")

	//act
	prom, err := promedio(notas...)

	//assert
	assert.Equal(t, expectedProm, prom)
	assert.Equal(t, expectedErr, err)
}
