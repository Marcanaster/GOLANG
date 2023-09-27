package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int16
}

func main() {

	var u usuario
	u.nome = "Marcos Reis"
	u.idade = 49
	fmt.Println(u)

	endereco := endereco{"Rua Manoel Vieiro", 201}

	usuario2 := usuario{"Marcos Reis", 21, endereco}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Marcos Reis"}
	fmt.Println(usuario3)
}
