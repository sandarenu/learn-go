package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Print("Hello Wold\n")
	fmt.Println(reflect.TypeOf("Test"))
	fmt.Println(reflect.TypeOf(true))
}
