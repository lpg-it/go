package main

import "fmt"

// 有两个变量 a 和 b，a 的值为 10，b 的值为 20，如何交换两个变量的值？
func main0201() {
	// 方法一：传统方式
	a, b := 10, 20
	temp := a
	a = b
	b = temp
	fmt.Println("a = ", a, "b = ", b)
}

func main(){
	// 方法二
	a, b := 10, 20
	a, b = b, a
	fmt.Println("a = ", a, "b = ", b)
}
