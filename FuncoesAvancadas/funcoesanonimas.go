package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Println(texto)
	}("Passando parametro sem return")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d",texto, 2)
	}("Passando parametro com return")

	fmt.Println(retorno)
}