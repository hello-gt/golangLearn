package main

import "fmt"

func main() {
	var num int
	var zs int

	for i := 1; i <= 100; i++ {
		if (i % 9) == 0 {
			zs += i
			num++
		}
	}
	fmt.Println(num)
	fmt.Println(zs)
}
