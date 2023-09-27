package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -10000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000000
	fmt.Println(numero2)
    //int32 = rune
	var numero3 rune = 100
	fmt.Println(numero3)
    //byte = int8
	var numero4 byte = 100
	fmt.Println(numero4)

	var numero5 float32 = 123.45
	fmt.Println(numero5)

	var numero6 float32 = 123.45
	fmt.Println(numero6)

	var str string = "Texto"
	fmt.Println(str)

	char := 'b'
	fmt.Println(char)

	var booleano bool  = true
	fmt.Println(booleano)

	var erro error = errors.New("Erro")
	fmt.Println(erro)
}