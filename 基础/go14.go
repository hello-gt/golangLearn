package main

import "fmt"

func main() {
	// var str string = "hello world"

	// for i := 0; i < len(str); i++ {
	// 	fmt.Println(str[i])
	// }

	str := "hello world"

	for key, val := range str {
		fmt.Println(key)
		fmt.Println(val)
		fmt.Println()
	}
}
