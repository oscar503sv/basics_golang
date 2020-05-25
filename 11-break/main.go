package main

import "fmt"

func main() {
	/*Las instrucciones break y continue*/
	//Break termina con el bucle
	//Continue deja ejecutar los codigos que siguen despues del continue
	for i := 0; i <= 10; i++ {
		if i == 5 {
			//fmt.Println("Instruccion break ejecutada")
			//break
			fmt.Print("Ejecución número 5\n")
			continue
		}
		fmt.Println(i)
	}
}
