package main

import "fmt"

func fibronacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibronacci(posicao - 2) + fibronacci(posicao - 1)
}

func main() {
	fmt.Println()

	posicao := uint(10)
	fmt.Println(fibronacci(posicao))

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibronacci(i))
	}
}
