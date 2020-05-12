package main

import "fmt"

func main() {
	// fmt.Print
	a, b := 10, 20
	fmt.Print(a)
	fmt.Print(b)

	// fmt.Println
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// fmt.Printf
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d", b)
}
