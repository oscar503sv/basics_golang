package main

import "fmt"

func main() {

	//saludo()

	nombre, edad := datos()

	fmt.Printf("Hola %s, este saludo es usando funciones con retorno \nTu edad es: %d", nombre, edad)

}

//Definiendo una función
func saludo() {
	fmt.Println("Hola, este es un saludo desde la función")

	//var nombre string

	//fmt.Print("Nombre: ")
	//fmt.Scanln(&nombre)
	//fmt.Printf("Hola %s, este es un saludo personalizado desde la función", nombre)
}

//funciones con retorno
func datos() (string, int) {
	var (
		nombre string
		edad   int
	)
	fmt.Print("Nombre: ")
	fmt.Scanln(&nombre)
	print("\nEdad: ")
	fmt.Scanln(&edad)

	return nombre, edad
}
