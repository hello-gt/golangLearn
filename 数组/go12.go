package main

import "fmt"

func main() {
	// var arr [3]float64
	// arr[0] = 3.1
	// arr[1] = 2.1
	// arr[2] = 54.1

	// data := 0.0
	// for i := 0; i < len(arr); i++ {
	// 	data += arr[i]
	// }
	// fmt.Println(data)

	var temp [3]int

	// fmt.Println("输入第一个数：")
	// fmt.Scanln(&temp[0])

	// fmt.Println("输入第2个数：")
	// fmt.Scanln(&temp[1])

	// fmt.Println("输入第3个数：")
	// fmt.Scanln(&temp[2])

	for i := 0; i < len(temp); i++ {
		fmt.Printf("输入第 %v 个数：", i)
		fmt.Scanln(&temp[i])
	}

	for i := 0; i < len(temp); i++ {
		fmt.Printf("输入第 %v 个数是 %v", i, temp[i])
	}

}
