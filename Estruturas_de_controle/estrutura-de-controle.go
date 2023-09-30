package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunga-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"

	}
}

//Aqui da a possibilidade de avaliar condiçoes diferentes
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunga-feira"
	case numero == 3:
		return "Terça-feira"
	case numero == 4:
		return "Quarta-feira"
	case numero == 5:
		return "Quinta-feira"
	case numero == 6:
		return "Sexta-feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func main() {
	fmt.Println()

	// if else

	numero := 10
	//no Go as chaves são obrigatórios
	if numero > 15 {
		fmt.Println("maior que 15")
	} else {
		fmt.Println("menor que 15")
	}

	if outronumero := numero; outronumero > 0 {
		fmt.Println("número maior que 0")
	} else {
		fmt.Println("número menor que zero")
	}

	if outronumero := -2; outronumero > 0 {
		fmt.Println("número maior que 0")
	} else if numero < -10 {
		fmt.Println("número menor que -10")
	} else {
		fmt.Println("número está 0 e -10")
	}

	// fim if else
	//switch

	dia := diaDaSemana(3)
	fmt.Println(dia)

	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)

	//fim switch

}
