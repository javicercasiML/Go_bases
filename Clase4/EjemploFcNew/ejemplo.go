package main

import (
	"errors"
	"fmt"
)

func main() {
	statusCode := 404
	if err := validateStatusCode(statusCode); err != nil {
		fmt.Printf("http request failed: %v", err)
		return
	}
	fmt.Println("the program ended successfully")
}
func validateStatusCode(code int) error {
	if code > 399 {
		return errors.New("unexpected http status code")
	}
	return nil
}
