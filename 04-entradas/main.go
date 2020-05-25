package main

import "fmt"

func main() {

	var (
		nombre   string
		apellido string
		edad     int
	)

	fmt.Print("***Ingrese sus datos***\n")
	// scanf para decirle que tipo de datos va recibir ejemplo:
	//fmt.Scanf("%s",&nombre)
	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&nombre)

	fmt.Print("Ingrese su apellido: ")
	fmt.Scanln(&apellido)

	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&edad)

	fmt.Printf("Nombre: %s %s \nEdad: %d", nombre, apellido, edad)

}
