package main

import "fmt"

func main() {
	function1()
}

func function1() {
	fmt.Printf("in 01")
	defer function2()
	fmt.Printf("in 02")
}
func function2() {
	fmt.Printf("in 03")
}
