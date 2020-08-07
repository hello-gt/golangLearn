package main

import "fmt"

func main() {
	var num int = 30

	for i := 1; i < num; i++ {
		if i == 29 {
			break
		}
		fmt.Println("hello world\n")
	}

}
