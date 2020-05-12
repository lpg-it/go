package main

import "fmt"

func main() {
	// fmt.Scanf
	var age int
	fmt.Println("请输入您的年龄：")
	fmt.Scanf("%d", &age)

	fmt.Println("您的年龄为：", age)

	// fmt.Scan
	var name string
	fmt.Println("请输入您的名字：")
	fmt.Scan(&name)

	fmt.Println("您的名字为：", name)
}
