package main

import (
	"errors"
	"fmt"
)

func main() {
	var w, h, a float64
	w = 5.2
	h = 3.5
	a = w * h / 10.0
	fmt.Println("area", a, "liters")
	fmt.Printf("area is: %T liters\n", a)
	fmt.Printf("area is: %0.3f liters\n", a)
	fmt.Printf("area is: %8.3f liters\n", a)
	printValue(a)
	repeat("This is a test", 3)
	a2 := calcArea(w, 3.33)
	printValue(a2)
	v := 10
	fmt.Printf("before v=%d\n", v)
	r := passValue(v)
	fmt.Printf("after v=%d, r=%d\n", v, r)
	a3 := calcArea(w, -3.33)
	printValue(a3)
	if a3 < 0 {
		err := fmt.Errorf("Incorrect value for area")
		fmt.Println(err.Error())
	}

	a4, e4 := calcAreaWithError(w, -3.33)
	printValue(a4)
	if e4 != nil {
		fmt.Println(e4.Error())
	}

	if x, e := calcAreaWithError(4, -5); e == nil {
		printValueWithMsg(x, "area calc with if assignment")
	} else {
		fmt.Println("Error", e.Error())
	}
}

func printValue(v float64) {
	fmt.Printf("area is: %8.3f liters\n", v)
}

func repeat(str string, times int) {

	for i := 0; i < times; i++ {
		fmt.Println(str)
	}
}

func calcArea(w float64, h float64) float64 {
	return w * h / 10.0
}

func calcAreaWithError(w float64, h float64) (area float64, err error) {
	a := calcArea(w, h)
	if a < 0 {
		return 0, errors.New("Incorrect value for area")
	}
	return a, nil
}

func passValue(v int) int {
	v = v * 2
	return v
}

func printValueWithMsg(v float64, msg string) {
	fmt.Printf("%s: %8.3f\n", msg, v)
}
