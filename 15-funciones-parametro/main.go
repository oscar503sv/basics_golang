package main

import "fmt"

func main() {

	var a, b int

	fmt.Print("Ingrese el dato 1: ")
	fmt.Scanln(&a)
	fmt.Print("Ingrese el dato 2: ")
	fmt.Scanln(&b)

	resultado := sumar(a, b)

	fmt.Printf("El resutado de la suma es: %d", resultado)

}

func sumar(a int, b int) int {
	suma := a + b
	return suma
}
