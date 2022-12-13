package main

import (
	"fmt"
)

func main() {
	var a = 100
	fmt.Println(double(a))
	var p1 = &a
	fmt.Println(p1)
	fmt.Println(*p1)
	*p1 = 300
	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Println("-------------")
	var b = 100
	var p2 = &b
	fmt.Println(*p2)
	fmt.Println(doubleP(p2))
	fmt.Println(*p2)
	fmt.Println(b)
}

func double(a int) int {
	a = a * 2
	return a
}

func doubleP(a *int) int {
	*a = *a * 2
	return *a
}
