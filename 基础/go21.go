package main

import "fmt"

func test(n1 int) {
	// n1 += 1

	if n1 > 2 {
		n1--
		test(n1)
	}

	fmt.Println("test() n1= ", n1)

}

func main() {
	n1 := 4
	test(n1)
	// fmt.Println("mian() n1=", n1)
}
