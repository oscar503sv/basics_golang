package main

import "fmt"

func main() {
	var (
		numero int
		salida string
	)

	fmt.Println("***INGRESE UN NUMERO DE 1 A 5***")
	fmt.Print("Ingrese el numero: ")
	fmt.Scanln(&numero)

	switch numero {
	case 1:
		salida = "Uno"
	case 2:
		salida = "Dos"
	case 3:
		salida = "Tres"
	case 4:
		salida = "Cuatro"
	case 5:
		salida = "Cinco"
	default:
		salida = "No está entre 1 y 5"
	}

	fmt.Printf("Valor ingresado: %s", salida)
}
