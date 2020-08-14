package main

import "fmt"

type myFunType func(int, int) int

func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	res := myFun2(getSum, 50, 60)
	fmt.Println("res2=", res)
}

//函数细节使用
