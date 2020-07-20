package main

import "fmt"

func main() {
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", num3s(2, 5, 6))
}

func num3s(a int, b int, c int) int {
	return a * b * c
}
