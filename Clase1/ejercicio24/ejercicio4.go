package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("Edad de Benjamin: ", employees["Benjamin"])
	i := 0

	for _, value := range employees {
		if value > 21 {
			i += 1
		}
	}

	fmt.Println("Empleados mayores de 21: ", i)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println("Lista final: ", employees)
}
