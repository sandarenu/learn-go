package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter your age: ")
	reader := bufio.NewReader(os.Stdin)
	age, _ := reader.ReadString('\n')
	fmt.Println("Your age is : ", age)
	agei, err := strconv.Atoi(strings.TrimSpace(age))

	if err != nil {
		log.Fatal("Error", err)
	}

	if agei > 50 {
		fmt.Println("You are getting older now...")
	} else if agei > 40 {
		fmt.Println("New era is begining...")
	} else {
		fmt.Println("You are still young...")
	}

}
