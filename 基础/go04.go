package main

import "fmt"

func main() {
	var i int32 = 100
	//var n1 float32 = i//错误
	var n1 float32 = float32(i)
	fmt.Println(n1)

}
