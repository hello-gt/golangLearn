package main

import (
	"fmt"
)

func main() {
	num := 100
	fmt.Printf("类型%T,值%v，地址%v", num, num, &num)
}
