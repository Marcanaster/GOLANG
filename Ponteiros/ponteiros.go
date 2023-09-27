package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	//não modifica variavel2
	fmt.Println(variavel1)

	//Ponteiro
	var variavel3 int
	//guarda enredeco de memoria
	var ponteiro *int

	variavel3 = 100
	//Vai dar erro
	//ponteiro = variavel3

	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)//mostra copia

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) //referencia de memoria 


}