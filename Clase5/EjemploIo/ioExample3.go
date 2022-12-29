package main

import (
	"io"
	"log"
	"os"
)

func _main() {

	var s = "Este texto es copiado hasta la ultima linea."
	// en terminal
	_, err := io.WriteString(os.Stdout, s)
	if err != nil {
		log.Fatal(err)
	}
}
