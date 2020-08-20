package main

import "fmt"

func main() {
	var arr [3]float64
	arr[0] = 3.1
	arr[1] = 2.1
	arr[2] = 54.1

	data := 0.0
	for i := 0; i < len(arr); i++ {
		data += arr[i]
	}
	fmt.Println(data)
}
