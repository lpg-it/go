package main

import "fmt"

func main0101() {
	// 使用 var 定义一个变量。注意：变量类型 要在 变量名 后面
	var hp int  // int 表示整数
	fmt.Println(hp)
}

func main0102(){
	//var hp int = 100 // 定义变量并初始化，等同于下面两行
	var hp int
	hp = 100
	fmt.Println(hp)
}

func main0103(){
	// 定义多个变量
	// 定义两个类型都是 int 的变量
	//var hp, mp int

	hp, mp := 100, 260
	fmt.Println(hp, mp)
}

func main(){
	// 匿名变量
	_, b := 3, 2  // b 为 2，丢弃 3
	fmt.Println(b)
}