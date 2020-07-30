package main

import (
	"fmt"
)

type B struct {
	thing int
}

func (this *B) change() { this.thing = 2 }

func (b B) write() string { return fmt.Sprint(b) }

func main() {
	var b1 B
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) //b2是指针
	b2.change()
	fmt.Println(b2.write())
}
