package main

import (
	"fmt"
)

type List []int

func (this List) Len() int        { return len(this) }
func (this *List) Append(val int) { *this = append(*this, val) }

func main() {
	//值
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)", lst, lst.Len())

	//指针
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)", plst, plst.Len())
}
