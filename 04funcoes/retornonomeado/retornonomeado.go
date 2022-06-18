package main

import "fmt"

func troca(p1, p2 int) (segundo int, primerio int) {
	segundo = p2
	primerio = p1
	return //return limpo
}

func main() {
	x, y := troca(1, 2)
	fmt.Println(x, y)
}
