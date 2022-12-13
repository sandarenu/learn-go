package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 1
	for i < 5 {
		fmt.Println("New loop", i)
		i++
	}
}
