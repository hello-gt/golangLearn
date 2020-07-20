package main

import "fmt"

var he int
var ji int
var cha int

func main() {
	var num1 int = 10
	var num2 int = 40

	he, ji, cha = yunsuan(num1, num2)
	PrintValues()

	he, ji, cha = yunsuan2(num1, num2)
	PrintValues()

}

func PrintValues() {
	fmt.Printf("num1 + num2 = %d, num1 x num2 = %d, num1 - num2 = %d\n", he, ji, cha)
}

func yunsuan(num1 int, num2 int) (int, int, int) {
	return num1 + num2, num1 * num2, num1 - num2
}

func yunsuan2(num1 int, num2 int) (x int, y int, z int) {
	x = num1 + num2
	y = num1 * num2
	z = num1 - num2
	return
}
