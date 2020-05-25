package main

import "fmt"

func main() {
	//Declaracion de variables
	var valor3 int
	var valor4 int

	fmt.Print("***Ingrese dos numeros enteros***\n")

	//Entrada
	fmt.Print("Ingrese el primer valor: ")
	fmt.Scanln(&valor3)
	fmt.Print("Ingrese el segundo valor: ")
	fmt.Scanln(&valor4)

	//Procesos
	fmt.Println("¿Son iguales?", valor3 == valor4)
	fmt.Println("¿Son diferentes?", valor3 != valor4)
	fmt.Println("¿El primer valor es mayor?", valor3 > valor4)
	fmt.Println("¿El segundo valor es mayor?", valor3 < valor4)
	fmt.Println("¿El Primero es menor o igual que el segundo?", valor3 <= valor4)
	fmt.Println("¿El Primero es mayor o igual que el segundo?", valor3 >= valor4)

}
