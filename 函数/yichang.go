package main

import "fmt"

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	num := 10
	num1 := 0
	num2 := num / num1
	fmt.Println(num2)
}
func main() {
	test()
	fmt.Println("继续。。。。")
}
