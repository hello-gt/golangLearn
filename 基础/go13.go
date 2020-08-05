package main

import "fmt"

func main() {
	var key byte

	fmt.Println("请输入：")
	fmt.Scanf("%c", &key)
	fmt.Println(key)
	switch key {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	default:
		fmt.Println("其他")
	}

}
