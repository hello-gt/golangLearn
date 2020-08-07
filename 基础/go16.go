package main

import (
	"fmt"
	"time"
)

func main() {

	temp := time.Now().Unix()

	fmt.Println(temp)

	var num int = 30

	for i := 1; i < num; i++ {
		if i == 29 {
			break
		}
		fmt.Println("hello world\n")
	}

}
