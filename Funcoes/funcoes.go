package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculoMatematico(n1, n2 int8)(int8, int8){
	soma := n1+n2
	subtracao := n1-n2

	return soma, subtracao

}

func main() {
	soma := somar(20, 20)
	fmt.Println(soma)

	var f = func(txt string) string{
		fmt.Println(txt)
		return txt

	}

	f("Texto1")

	resultado := f("Text2")
	fmt.Println(resultado)

	resultadosSoma, resultadoSubtracao := calculoMatematico(10, 15)
	fmt.Println(resultadosSoma, resultadoSubtracao )

	resultadosSoma2, _ := calculoMatematico(10, 15)
	fmt.Println(resultadosSoma2)
}