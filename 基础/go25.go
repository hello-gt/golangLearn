package main

import "fmt"

func main() {
	f := test()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}

func test() func(int) int {
	var temp int = 10

	return func(n int) int {
		temp += n
		return temp
	}

}

// function test(n) {
// 	var temp = 10

// 	return function (n) {
// 		temp += n
// 		return temp
// 	}
// }

// function main(){
// 	console.log(test(1))
// 	console.log(test(2))
// 	console.log(test(3))
// }
