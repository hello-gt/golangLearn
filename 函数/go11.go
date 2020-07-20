package main

import "fmt"

func main() {
	doDbperations()
}

func connectToDb() {
	fmt.Println("链接数据库")
}
func closeToDb() {
	fmt.Println("最后一步关闭数据库")
}

func doDbperations() {
	connectToDb()
	fmt.Println("打印数据库")
	defer closeToDb()
	fmt.Println("数据库资源")
}

//defer 的使用  最后一步释放资源
