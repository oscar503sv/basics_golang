package main

import (
	"fmt"
	"strings"
)

func main() {

	cadena := "Hola, esta es una cadena string, Hola Golang"

	fmt.Println(strings.ToUpper(cadena))
	fmt.Println(strings.ToLower(cadena))
	fmt.Println(strings.Replace(cadena, "Hola", "Hello", 1))
	fmt.Println(strings.Replace(cadena, "Hola", "Hello", 2))

	fmt.Println(strings.Split(cadena, " ")) //convirtiendo a array

	var slicen1 []string

	slicen1 = append(strings.Split(cadena, " ")) //convirtiendo a slicen

	fmt.Println(slicen1)
	fmt.Println(slicen1, len(slicen1))

}
