package main

import "fmt"

func main() {
	var num1 int = 95
	var num2 int = 85
	var min, max int
	min, max = getRes(num1, num2)
	fmt.Printf("this is min：%d,this is max: %d", min, max)
}
func getRes(num1 int, num2 int) (min int, max int) {
	if num1 > num2 {
		min = num2
		max = num1
	} else {
		min = num1
		max = num2
	}
	return
}

//最大值和最小值判断
