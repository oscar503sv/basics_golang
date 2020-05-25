package main

import "fmt"

func main() {

	//Operacion en asignacion
	a := 2
	a += 3 // Sumando el valor que ya tiene a + 3

	fmt.Println(a)
	// BUCLES
	/* En Go no existe el bucle while y do while simplemente existe
	el bucle for que se puede utilizar para que se comporte como un while*/
	c := 0

	//Bucle for comportandose como un while
	for c <= 5 {
		fmt.Println(c)
		c++
	}

	//Bucle for tradicional
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}
