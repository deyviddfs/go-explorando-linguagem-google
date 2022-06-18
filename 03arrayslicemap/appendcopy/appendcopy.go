package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	//array1 = append(array1, 4, 5, 6) // Não funciona, não é possível fazer append de array
	slice1 = append(slice1, 4, 5, 6) //append aumenta automaticamente tamanho da array do slice
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // copy não aumenta automáticamente o tamanho do array do slice
	fmt.Println(slice2)
}
