package main

import (
	"fmt"
	"time"
)

func notaParaConceito(n float64) string {
	nota := int(n)
	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota <= 9:
		return "B"
	case nota >= 5 && nota <= 8:
		return "C"
	case n >= 3 && nota <= 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	t := time.Now()
	switch { //switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}

	fmt.Println(notaParaConceito(9.1))
	fmt.Println(notaParaConceito(8.2))
	fmt.Println(notaParaConceito(6.3))
	fmt.Println(notaParaConceito(3.4))
	fmt.Println(notaParaConceito(1.8))
}
