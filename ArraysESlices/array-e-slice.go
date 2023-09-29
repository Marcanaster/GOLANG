package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println()
	var array1 [5]int
	fmt.Println(array1)

	var array2 [5]string
	fmt.Println(array2)

	array3 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array3)

	array4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array4)

	//Não rpecisa especificar o tamnho
	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array4))

	slice1 = append(slice1, 21)
	fmt.Println(slice1)
	
	slice2 := array3[1:3]
	fmt.Println(slice2)

	array3[1] = "Posicao alterada"
	fmt.Println(slice2)

	

}
