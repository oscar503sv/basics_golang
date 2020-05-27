package main

import "fmt"

func main() {

	var array1 [3]string
	array2 := [3]int{1, 2, 3} // declarando array ya con sus valores inicializados

	array1[0] = "Oscar"
	array1[1] = "Mauricio"
	array1[2] = "El Salvador"

	fmt.Printf("Nombre: %s %s \nPaís: %s\n", array1[0], array1[1], array1[2])
	fmt.Println(array2)

}
