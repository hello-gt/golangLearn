package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newS string
	fmt.Printf("the size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("the integer is:%d\n", an)

	an += 5

	newS = strconv.Itoa(an)
	fmt.Printf("the new string is :%s\n", newS)
}

//字符串类型转换
