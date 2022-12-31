package main

import (
	"fmt"

	"desafio/internal/tickets"
)

const (
	earlymorning = "madrugada"
	morning      = "mañana"
	afternoon    = "tarde"
	evening      = "noche"
	path         = "/Users/jcercasi/desafio-go-bases/tickets.csv"
)

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Ejecución Finalizada.\n ")
	}()

	err := tickets.Reader(path)
	if err != nil {
		panic(err)
	}

	country := "Brazil"
	total, err := tickets.GetTotalTickets(country)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nViajan a %s %d personas\n\n", country, total)

	turns := tickets.GetMornings()
	fmt.Printf("Viajan %d de %s, %d de %s, %d de %s y %d de %s\n\n", turns[0], earlymorning, turns[1], morning, turns[2], afternoon, turns[3], evening)

	cant, err := tickets.AverageDestination(country)
	if err != nil {
		panic(err)
	}
	fmt.Printf("El promedio de personas que viajan a %s es: %.2f%%\n\n", country, cant)
}
