package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImpuestoA(t *testing.T) {
	//arange
	var sueldo uint = 40000
	expectedImpuesto := "0%"

	//act
	impuesto := impuesto(sueldo)

	//assert
	assert.Equal(t, expectedImpuesto, impuesto)
}

func TestImpuestoB(t *testing.T) {
	//arange
	var sueldo uint = 52000
	expectedImpuesto := "17%"

	//act
	impuesto := impuesto(sueldo)

	//assert
	assert.Equal(t, expectedImpuesto, impuesto)

}

func TestImpuestoC(t *testing.T) {
	//arange
	var sueldo uint = 160000
	expectedImpuesto := "27%"

	//act
	impuesto := impuesto(sueldo)

	//assert
	assert.Equal(t, expectedImpuesto, impuesto)

}
