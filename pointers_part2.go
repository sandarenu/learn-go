package main

import "fmt"

type Vehicle struct {
	Class string
	Brand string
}

// Constructor Pattern.
// Use factory method to create new instance
func NewVehicle(c string, b string) *Vehicle {
	return &Vehicle{Class: c, Brand: b}
}

type Adder struct {
	start int
}

func (a *Adder) Add(i int) *Adder {
	//Supporting nil receiver
	if a == nil {
		return &Adder{i}
	}
	a.start += i
	return a
}

type Score struct {
	Value int
}

type HighestScore Score
type LowestScore Score

func main() {

	//This is a nil pointer
	var v1 *Vehicle
	fmt.Printf("nil pointer %p\n", v1)
	v1 = NewVehicle("Car", "Toyota")
	fmt.Printf("value assigned pointer %p\n", v1)

	a := &Adder{0}
	fmt.Println(a)
	a.Add(5)
	a.Add(5)
	a.Add(5)
	fmt.Println(a)

	var a1 *Adder
	a1 = a1.Add(5)
	fmt.Println(a1)

	//assign function to variable
	//method value
	f1 := a1.Add
	f1(10)
	f1(10)
	f1(10)
	fmt.Println(a1)

	//method expression
	f2 := (*Adder).Add
	f2(a1, 20)
	fmt.Println(a1)

	a2 := f2(nil, 10)
	fmt.Println(a2)

	s1 := Score{56}
	fmt.Println("Score", s1)
	h1 := HighestScore{98}
	fmt.Println("HiScore", h1)
	s1 = Score(h1)
	fmt.Println("Score", s1)

}
