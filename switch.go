package main

import (
    "math/rand"
    "fmt"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)
    switch i {
    case 0:
        fmt.Println("Zero")
    case 1:
        fmt.Println("One")
    default:
        fmt.Println("I'm lazy to type")
    }

    switch sum := i + 10; sum {
    case 10:
        fmt.Println("Sum is 10", sum)
    default:
        fmt.Println("Sum is an other number", sum)
    }

    // Blank Switch, allows us to do comparisons
    switch sum := i + 10;  {
    case sum < 10:
        fmt.Println("Sum is less than 10", sum)
    default:
        fmt.Println("Sum is greater than 10", sum)
    }
}