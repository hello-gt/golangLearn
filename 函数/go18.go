package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

func Adder() func(a int) int {
	var b int
	return func(a int) int {
		b += a
		return b
	}
}

//匿名函数
