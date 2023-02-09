package main

import (
	"fmt"
)

type AppStatus int

const (
	Draft AppStatus = iota
	Pending
	LimitedProduction
	ActiveProduction
	Rejected
)

func main() {

	s := ActiveProduction
	fmt.Printf("%v\n", s)
	s = Rejected
	fmt.Printf("%v\n", s)

}
