package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá mundo !")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("marcnaastermail.com")
	fmt.Println(erro)
}
