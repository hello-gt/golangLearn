package main

import "fmt"

func main() {
	fv := func() { fmt.Println("hello world") }
	fv()
	f()
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

//闭包  即  匿名函数
