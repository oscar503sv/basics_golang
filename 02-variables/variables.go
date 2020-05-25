package main

import "fmt"

func main() {
	//Definir variables
	var nombre string
	var a, b, c int
	const pi = 3.141592 //Declaracion de una constante
	nombre = "Oscar"

	fmt.Println(nombre)

	a = 2
	b = 4
	c = 5

	var resultado = a + b + c

	fmt.Println("Resultado:", resultado)

	var (
		booleano bool
		apellido = "Aragon"
		edad     = 27
	)

	booleano = true
	anio := 2020 //Declaracion corta de variables en go, hay que asignarles el valor

	println("Pi:", pi)
	println("Bool:", booleano)
	println("Mi nombre es:", nombre+" "+apellido)
	println("Mi edad es:", edad)
	println("Estamos en el año:", anio)
}
