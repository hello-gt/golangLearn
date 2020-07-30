package main

import (
	"fmt"
	"time"
)

type MyTime struct {
	time.Time
}

func (this MyTime) first3Chars() string {
	return this.Time.String()[0:4]
}
func main() {
	m := MyTime{time.Now()}
	//调用匿名Time上的String
	fmt.Println("Full time now:", m.String())
	//调用myTime。first3Chars
	fmt.Println("First 3 chars:", m.first3Chars())

}
