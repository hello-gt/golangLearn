package main

import (
	"fmt"
)

func main() {
	var num int = 100
	var add int
	var dangqian int = 20
	for i := 1; i < num; i++ {
		if add > dangqian {
			break
		}
		add += i

		fmt.Println(i)

	}

}
