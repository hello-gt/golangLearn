package main

import "fmt"

func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Printf("zhi wei:", *reply)
}
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

//函数改变外部变量 通过指针
