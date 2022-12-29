package main

import (
	"errors"
	"fmt"
)

func sayHi(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}
	return fmt.Sprintf("Hi, %s .", name), nil
}

func main() {
	greeting, err := sayHi("Javi")
	fmt.Println(greeting, err)

}
