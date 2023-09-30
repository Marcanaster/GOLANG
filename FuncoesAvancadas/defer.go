package main

import "fmt"

func funcao1() {
	fmt.Println("excutando função 1")
}

func funcao2() {
	fmt.Println("excutando função 2")
}

func alunoestaaprovado(nota1, nota2 float32) bool{
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	media := (nota1 + nota2) / 2
	
	if media >= 6{
		defer fmt.Println("Media calculada. Resultado será retornado")
		return true
	}
	return false

}
func main() {
	//defer funcao1()
	funcao2()

	fmt.Println(alunoestaaprovado(7, 8))
}
