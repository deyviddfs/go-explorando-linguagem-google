package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, valor := range numeros {
		total += valor
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(1, 2, 3, 4, 4))
}
