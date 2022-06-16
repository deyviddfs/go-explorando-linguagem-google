package main

import "fmt"

func main() {
	x := 1
	y := 2

	//apenas postfix
	x++ // x = x + 1
	fmt.Println(x)

	y-- // y = y -1
	fmt.Println(y)

	//Não aceita e não existe --y (como no Java)
	//fmt.Println(x == y--)
}
