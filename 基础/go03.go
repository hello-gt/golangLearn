package main

import "fmt"

func main() {
	var a string = "hello"
	fmt.Println(a)

	var b string = `hello \n`
	var c string = `world`
	var d string = b + c
	fmt.Println(d)
}
