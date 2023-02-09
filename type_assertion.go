package main

import ("fmt"
    "log"
)
type MyInt int
//Type assertions can only be applied to interface types and are checked at runtime
func main(){

    var i interface{}
    var myInt1 MyInt = 10
    i = myInt1
    //Do not compile, i is not an int
    //j := i + 10
    // i is asserted to MyInt
    j := i.(MyInt) + 20
    fmt.Println(j)

    //Handling assertion errors with "ok pattern"
    //This code compiles since assertion is validated at runtime.
    p, ok := i.(MyInt)
    if !ok {
       log.Fatal("i can't be converted to int")
    }
    fmt.Println("p", p)

    printMessage(10)
    printMessage("Hi")
    printMessage(i)
}

func printMessage(i interface{}){
    switch  j := i.(type) {
    case nil:
        fmt.Println("It is nil")
    case int:
        fmt.Println("It is an int", j)
    case MyInt:
        fmt.Println("It is a MyInt", j)
    default:
        fmt.Println("Unknow type", j)
    }
}