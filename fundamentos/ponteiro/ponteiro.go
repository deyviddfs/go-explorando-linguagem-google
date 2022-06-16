package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i //pegando o endreço da variável
	*p++   // desreferenciando (pegando o valor)
	i++

	// Go não tem aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
