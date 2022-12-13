package main

import "fmt"

func main() {
	var numbers [5]int
	numbers[0] = 10
	fmt.Println("Array: ", numbers)
	for i := 0; i < 5; i++ {
		numbers[i] = i * 100
	}
	for i := 0; i < 5; i++ {
		fmt.Println(numbers[i])
	}

	a1 := []int{1, 2, 3, 4}
	fmt.Println(a1)

	a2 := [10]int{1, 2, 3, 4}
	fmt.Println(a2)

	a3 := []string{
		"This",
		"is",
		"a",
		"test",
	}
	fmt.Println(a3)
	fmt.Printf("message %#v\n", a3)

	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}

	for i, v := range a3 {
		fmt.Printf("index %d - %v\n", i, v)
	}

	for _, v := range a3 {
		fmt.Println(v)
	}
}
