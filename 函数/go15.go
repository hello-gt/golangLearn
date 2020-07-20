package main

import "fmt"

func main() {
	callback(1, Add)
}

func Add(a int, b int) {
	fmt.Printf("the sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

//函数当成参数
