package main

import "fmt"

type myFunType func(int, int) int

func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func getSum3(num1 int, num2 int) (res int, res2 int) {
	res = num1 + num2
	res2 = num1 - num2
	return res, res2
}

func main() {
	res, res2 := getSum3(50, 60)
	fmt.Println("res=", res)
	fmt.Println("res2=", res2)
}

//函数细节使用
