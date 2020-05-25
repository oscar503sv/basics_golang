package main

import "fmt"

func main() {

	//Declarando la variable
	var valor int

	fmt.Print("***INGRESE UN NUMERO ENTERO***\n")
	fmt.Print("Ingrese el numero: ")
	fmt.Scanln(&valor)

	if valor == 0 {
		fmt.Printf("El numero %d es neutro", valor)
	} else if (valor % 2) == 0 {
		fmt.Printf("El numero %d es par", valor)
	} else {
		fmt.Printf("El numero %d es impar", valor)
	}
}
