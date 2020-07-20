package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("time: %s\n", delta)
}
