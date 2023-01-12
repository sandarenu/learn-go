package main

import (
	"fmt"
)

func main() {
	var s1 []string
	fmt.Println("s1", s1)
	s1 = make([]string, 7)
	s1[0] = "a"
	s1[1] = "a"

	fmt.Println("s1 after", s1)
	s2 := make([]int, 10)
	fmt.Println("s2", s2)

	a1 := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println("a1", a1)
	s3 := a1[0:3]
	fmt.Println("s3", s3)
	a1[0] = "A"
	fmt.Println("a1", a1)
	fmt.Println("s3", s3)

	s4 := make([]int, 4)
	s4[0] = 1
	s4[1] = 2
	fmt.Println("s4 ", s4)
	s5 := append(s4, 3)
	fmt.Println("s4 after", s4)
	fmt.Println("s5", s5)
}
