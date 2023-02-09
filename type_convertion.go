package main

import "fmt"

func main() {
	var a int = 100
	var b float64 = 2.45

	fmt.Println(float64(a) * b)
	fmt.Println(a * int(b))
}
