package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minimunFunction(nums ...float64) (result float64) {
	result = nums[0]
	for _, num := range nums {
		if num < result {
			result = num
		}
	}
	return
}

func averageFunction(nums ...float64) (result float64) {
	var sum float64

	for _, num := range nums {
		sum += num
	}
	result = sum / float64(len(nums))
	return
}
func maximunFuntion(nums ...float64) (result float64) {
	for _, num := range nums {
		if result < num {
			result = num
		}
	}
	return
}

type OperationFunc func(nums ...float64) float64

func operation(oper string) (fn OperationFunc, err error) {
	switch oper {
	case minimum:
		return minimunFunction, err
	case maximum:
		return maximunFuntion, err
	case average:
		return averageFunction, err
	default:
		err = errors.New("Operacion invalida")
		return nil, err
	}
}
func main() {

	operaciones := []string{minimum, average, maximum}

	for _, oper := range operaciones {
		fun, err := operation(oper)

		if err != nil {
			fmt.Println(err, " en: ", oper)
		} else {
			value := fun(2, 3, 3, 4, 10, 2, 4, 5)
			fmt.Println("\nEl numero", oper, " es: ", value)
		}
	}
	//  minFunction, err := operation(minimum)
	//  averageFunc, err := operation(average)
	//  maxFunc, err := operation(maximum)

	//  minValue := minFunction(2, 3, 3, 4, 10, 2, 4, 5)
	//  averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	//  maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

}
