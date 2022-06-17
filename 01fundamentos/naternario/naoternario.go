package main

import "fmt"

// Não tem operador ternário
func obterResulatdo(nota float64) string {
	// (nota >= 6) ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResulatdo(6.2))
}
