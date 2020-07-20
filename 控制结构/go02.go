package main

import "fmt"

func main() {
	var num1 int = 100

	switch num1 {
	case 98, 99:
		fmt.Println("It is 99,98")
	case 100:
		fmt.Println("It is 100")
	default:
		fmt.Println("not select")
	}
}

// switch case
