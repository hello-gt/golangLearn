package main

import "fmt"

func main() {
	var day int = 97
	var week int = day / 7
	var num int = day % 7
	fmt.Println(week)
	fmt.Println(num)
}
