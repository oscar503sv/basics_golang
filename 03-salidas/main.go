package main

import "fmt"

func main() {
	fmt.Print("Hola mundo\n") //el print no hace saltos de linea para hacer saltos de linea \n
	fmt.Println("Hola mundo") //println si hace un salto de linea

	nombre := "Oscar"
	edad := 27
	pi := 3.141592

	//Concatenacion de datos
	fmt.Println("Nombre:", nombre, "Edad:", edad, "\nValor de pi:", pi)

	//formateando los datos
	// %s para mostrar tipos string
	// %d para mostrar tipos enteros
	// %f para mostrar tipos flotantes
	fmt.Printf("Nombre: %s Edad: %d \nValor de pi: %f", nombre, edad, pi)
}
