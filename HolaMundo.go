package main

import "fmt"

//Hola Mundo
/**Hola Mundo*/

func main() {
	//fmt.Println("Hola Mundo")
	var numero int
	numero = 25
	fmt.Println(numero)
	numero = 40
	fmt.Println(numero)

	nombre := "Alejandro"
	fmt.Println(nombre)

	nombre, numero = "Eilen", 25
	fmt.Println(nombre, numero)

	nombre2 := "Morales"
	fmt.Println(nombre, nombre2)
	nombre, nombre2 = nombre2, nombre
	fmt.Println(nombre, nombre2)
}
