package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println()

	//for

	i := 0
	fmt.Println(i)
	for i < 10{
		i++
		fmt.Println("incrementando i", i)
		time.Sleep(time.Second)
	}

	fmt.Println(i)
	for j := 0; j < 10; j++{
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}
	//fim for

	//range

	nomes := [3]string{"marcos", "Marcelo", "Marcio"}

	for indice, nome := range nomes{
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes{
		fmt.Println(nome)
	}

	//fim range
}