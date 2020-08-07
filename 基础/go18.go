package main

import (
	"fmt"
)

func main() {

	var num int = 3

	for i := 1; i <= num; i++ {

		var usrname string
		var paw string

		fmt.Println("请输入用户名：")
		fmt.Scanln(&usrname)
		fmt.Println("请输入密码：")
		fmt.Scanln(&paw)

		if usrname == "zwj" && paw == "888" {
			fmt.Println("成功")
			break
		} else {
			fmt.Println("失败")
		}

		fmt.Printf("你还有%v次机会\n", num-i)
	}

}
