package main

import "fmt"

func inverterSinal(numero int) int{
	return numero * -1
}

func inverterSinalComPonteiro(numero *int){
	*numero = *numero * -1
}

func main() {

	numero := 20
	numeroInvetido := inverterSinal(numero)
	fmt.Println(numeroInvetido)
	fmt.Println(numero)

	novoNumero := 40
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}