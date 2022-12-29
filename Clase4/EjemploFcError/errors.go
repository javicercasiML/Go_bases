package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// Unwrap()
	error_1 := fmt.Errorf("first error")
	error_2 := fmt.Errorf("second error: %w", error_1)

	err := errors.Unwrap(error_2)
	if err == error_1 {
		fmt.Println("same error")
	}

	//Is()
	_, err1 := os.Open("not_exist.txt")
	var notExist error = fs.ErrNotExist
	if errors.Is(err1, notExist) {
		fmt.Println("The file does not exist")
	}

	//As()
	_, err3 := os.Open("not_exist.txt")
	var pathError *fs.PathError
	if errors.As(err3, &pathError) {
		fmt.Println("Path error", pathError.Path)
	}

}
