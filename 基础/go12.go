package main

import "fmt"

func main() {
	var a float32
	var b int
	var c bool

	fmt.Println("请输入升高:")
	fmt.Scanln(&a)

	fmt.Println("请输入薪水:")
	fmt.Scanln(&b)

	fmt.Println("是否帅:")
	fmt.Scanln(&c)

	if a > 180 && b > 10000 && c == true {
		fmt.Println("嫁给他")
	} else if a > 180 || b > 10000 || c == true {
		fmt.Println("算了")
	} else {
		fmt.Println("不假")
	}
}
