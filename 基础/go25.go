package main

import "fmt"

var name string = "jack"

func main() {
	fmt.Println(name)
	test1()
	test2()
	test3()
}

func test1() {
	fmt.Println(name)
}

func test2() {
	name := "tom"
	// name := "tom"
	fmt.Println(name)

}

func test3() {
	fmt.Println(name)
}
