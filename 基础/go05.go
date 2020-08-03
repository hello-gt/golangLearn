package main

import "fmt"

func main() {
	var num int = 9
	fmt.Printf("num address=%v\n", &num)

	var ptr *int
	ptr = &num //指针只能接受地址

	*ptr = 10
	fmt.Println("num =", num)
}
