package main

import "fmt"

type TZ int

func main() {
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf("c has ths value: %d", c)
	s := "hello"
	s += " world"
	fmt.Println(s)

}
