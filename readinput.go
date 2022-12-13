package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter your age: ")
	reader := bufio.NewReader(os.Stdin)
	age, _ := reader.ReadString('\n')
	fmt.Println("Your age is : ", age)
}
