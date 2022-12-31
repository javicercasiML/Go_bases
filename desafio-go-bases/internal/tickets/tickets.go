package tickets

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	id      string
	nombre  string
	email   string
	destino string
	hora    string
	precio  string
}

var Customers = []Ticket{}

var (
	ErrOpen = errors.New("\nError al abrir el archivo.")
	ErrScan = errors.New("\nError al leer el archivo.")
	ErrDest = errors.New("\nError: Destino no encontrado")
)

// ejemplo 1
func GetTotalTickets(destination string) (cant int, err error) {
	for _, cli := range Customers {
		if cli.destino == destination {
			cant++
		}
	}
	if cant == 0 {
		return 0, ErrDest
	}
	return cant, nil
}

// ejemplo 2
func GetMornings() (turns []int) {

	earlymorning, morning, afternoon, evening := 0, 0, 0, 0
	for _, cli := range Customers {
		hora, _ := strconv.Atoi(strings.Split(cli.hora, ":")[0])
		switch {
		case hora >= 0 && hora <= 6:
			earlymorning++
		case hora >= 7 && hora <= 12:
			morning++
		case hora >= 13 && hora <= 19:
			afternoon++
		case hora >= 20 && hora <= 23:
			evening++
		}
	}
	turns = append(turns, earlymorning, morning, afternoon, evening)
	return
}

// ejemplo 3
func AverageDestination(destination string) (cant float64, err error) {
	total, err := GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	return (float64(total) / float64(len(Customers))) * 100, nil
}

func Reader(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return ErrOpen
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		col := strings.Split(scanner.Text(), ",")
		ticket := Ticket{col[0], col[1], col[2], col[3], col[4], col[5]}
		Customers = append(Customers, ticket)
	}

	if err := scanner.Err(); err != nil {
		return ErrScan
	}
	return
}
