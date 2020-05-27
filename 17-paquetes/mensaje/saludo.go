package mensaje

import "fmt"

//Las funciones privadas empiezan con letra miniscula
func funcionPrivada() {
	fmt.Println("Hola, soy una función privada")
}

//Las funciones públicas se deben comentar como el ejemplo de abajo sino el compilador muestra error

//FuncionPublica esta es una funcion publica porque empieza con mayuscula
func FuncionPublica() {
	fmt.Println("Hola, soy una función pública")

}
