package tickets

import (
	"errors"
	"testing"

	// go get github.com/stretchr/testify/assert
	"github.com/stretchr/testify/assert"
)

func TestReaderOk(t *testing.T) {
	//arange
	path := "/Users/jcercasi/desafio-go-bases/tickets.csv"

	//act
	err := Reader(path)

	//assert
	assert.Equal(t, nil, err)
}

func TestReaderFail(t *testing.T) {
	//arange
	path := "/Users/jcercasi/desafio-go-bases/tickets.cs"
	ErrOpen = errors.New("\nError al abrir el archivo.")

	//act
	err := Reader(path)

	//assert
	assert.Equal(t, ErrOpen, err)
}

func TestGetTotalTickets(t *testing.T) {

	country1, country2 := "Argentina", "Brazil"
	ExpectedTotal1, ExpectedTotal2 := 15, 45

	//act
	total1, err1 := GetTotalTickets(country1)
	total2, err2 := GetTotalTickets(country2)

	//assert
	assert.Equal(t, ExpectedTotal1, total1)
	assert.Nil(t, err1)

	assert.Equal(t, ExpectedTotal2, total2)
	assert.Nil(t, err2)

}

func TestGetTotalErr(t *testing.T) {

	country, country1, country2 := "Argtina", "Brasil", "Chhina"
	ErrExpected := ErrDest

	//act
	_, err := GetTotalTickets(country)
	_, err1 := GetTotalTickets(country1)
	_, err2 := GetTotalTickets(country2)

	//assert
	assert.Equal(t, ErrExpected, err)
	assert.Equal(t, ErrExpected, err1)
	assert.Equal(t, ErrExpected, err2)

}

func TestGetMornings(t *testing.T) {

	ExpectedEar, ExpectedMor, ExpectedAft, ExpectedEve := 304, 256, 289, 151

	//act
	turnos := GetMornings()

	//assert
	assert.Equal(t, ExpectedEar, turnos[0])
	assert.Equal(t, ExpectedMor, turnos[1])
	assert.Equal(t, ExpectedAft, turnos[2])
	assert.Equal(t, ExpectedEve, turnos[3])

}

func TestAverageDestination(t *testing.T) {

	country, country1, country2 := "Argentina", "Brazil", "China"
	ExpectedAvg, ExpectedAvg1, ExpectedAvg2 := 1.50, 4.50, 17.8

	//act
	average, err := AverageDestination(country)
	average1, _ := AverageDestination(country1)
	average2, _ := AverageDestination(country2)

	//assert
	assert.Equal(t, ExpectedAvg, average)
	assert.Equal(t, ExpectedAvg1, average1)
	assert.Equal(t, ExpectedAvg2, average2)
	assert.Nil(t, err)

}

func TestAverageDestFail(t *testing.T) {

	country, country1, country2 := "Argtina", "Brasil", "Chhina"
	ExpectedErr := ErrDest

	//act
	total, err := AverageDestination(country)
	_, err1 := AverageDestination(country1)
	_, err2 := AverageDestination(country2)

	//assert
	assert.Equal(t, 0.0, total)
	assert.Equal(t, ExpectedErr, err)
	assert.Equal(t, ExpectedErr, err1)
	assert.Equal(t, ExpectedErr, err2)

}
