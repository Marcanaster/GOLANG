package main

import "fmt"

type usuario struct {
	nome  string
	idade uint16
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s ", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func main() {
	usuario1 := usuario{"Marcos", 50}
	usuario1.salvar()

	maiordeIdade := usuario.maiorDeIdade(usuario1)
	fmt.Println(maiordeIdade)
}
