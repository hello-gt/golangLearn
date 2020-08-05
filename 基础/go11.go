package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 50
	c := 5
	sumAdd(a, b, c)
}

func sumAdd(a int, b int, c int) {
	if (b*b - 4*a*c) > 0 {
		fmt.Println("两个解")
	} else if (b*2 - 4*a*c) == 0 {
		fmt.Println("一个解")
	} else {
		fmt.Println("无解")
	}
}
