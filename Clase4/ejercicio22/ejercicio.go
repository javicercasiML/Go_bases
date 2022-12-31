package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lectura("customers.txt")
}

func lectura(path string) {
	//var lineas []string
	var lineas string

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Ejecución Finalizada.\n ")
	}()

	file, err := os.Open(path)
	if err != nil {
		panic("\nEl archivo indicado no fue encontrado o está dañado")
	}

	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)
	for fileScan.Scan() {
		lineas += "\n" + fileScan.Text()
		//lineas = append(lineas, fileScan.Text())
	}

	fmt.Println("\nDatos del archivos:", lineas, "\n ")
	file.Close()

}
