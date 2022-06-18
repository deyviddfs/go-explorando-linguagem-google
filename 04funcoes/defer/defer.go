package main

import "fmt"

func obterValorAprovadoDefer(numero int) int {
	defer fmt.Print(" fim! ")
	if numero > 5000 {
		fmt.Print(" Valor alto... ")
		return 5000
	} else {
		fmt.Print(" Valor baixo... ")
		return numero
	}
}

func obterValorAprovado(numero int) int {
	fmt.Print(" fim! ")
	if numero > 5000 {
		fmt.Print(" Valor alto... ")
		return 5000
	} else {
		fmt.Print(" Valor baixo... ")
		return numero
	}
}

func main() {
	fmt.Println("Defer")
	fmt.Println(obterValorAprovadoDefer(6000))
	fmt.Println(obterValorAprovadoDefer(3000))
	fmt.Println("------------------------")
	fmt.Println(" Sem Defer")
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
