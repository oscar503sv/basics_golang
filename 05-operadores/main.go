package main

import "fmt"

func main() {
	//Operadores aritmeticos

	var valor1, valor2 int

	fmt.Print("***Ingrese dos numeros enteros***\n")

	//entrada
	fmt.Print("Ingrese el primer valor: ")
	fmt.Scanln(&valor1)
	fmt.Print("Ingrese el segundo valor: ")
	fmt.Scanln(&valor2)

	//procesos aritmeticos
	suma := valor1 + valor2
	resta := valor1 - valor2
	multiplicacion := valor1 * valor2
	division := valor1 / valor1
	modulo := valor1 % valor2

	fmt.Printf("La suma es: %d \nLa resta es: %d \nLa multiplicacion es: %d \nLa division es: %d \nEl modulo es: %d", suma, resta, multiplicacion, division, modulo)
}
