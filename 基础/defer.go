package main

import "fmt"

func main() {
	result := test(10, 20)
	fmt.Println("result=", result)
}

func test(num1 int, num2 int) int {
	defer fmt.Println("num1=", num1)
	defer fmt.Println("num2=", num2)

	res := num1 + num2
	fmt.Println("res=", res)
	return res
}
