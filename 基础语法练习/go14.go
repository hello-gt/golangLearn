package main

import "fmt"

func main() {
	var i1 = 15
	fmt.Printf("An integer: %d, its location in memory :%p\n", i1, &i1)

	var pi *int
	pi = &i1

	fmt.Printf("pis location in memory :%p\n", pi)
}

//指针
