package main

import "fmt"

func main() {
	var arr [26]byte

	//联系1
	// temp := 'A'

	var temp byte = 'A'

	for i := 0; i < len(arr); i++ {
		arr[i] = byte(temp) + byte(i)
	}

	// for _, value := range arr {
	// 	fmt.Printf("%c\n", value)
	// }

	var max byte
	var xiabiao int
	var res int
	for index, value := range arr {
		if value > max {
			max = value
			xiabiao = index
		}
		res += int(value)
	}
	fmt.Printf("%c\n", max)
	// fmt.Println(max)
	fmt.Println(xiabiao)
	fmt.Println(res)
	// var jun float64
	jun := float64(res) / float64(len(arr))
	fmt.Println(jun)

}
