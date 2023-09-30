package main

import "fmt"

func main() {

	fmt.Println("")

	usuario := map[string]string{
		"nome": "Marcos",
		"sobrenome": "Reis",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome":{
			"primeiro": "Marcos",
			"ultimo": "Reis",
		},
		"curso":{
			"nome": "Engenharia",
			"campus": "campus 1",
		},

	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Sagitário",
	}

	fmt.Println(usuario2)
}