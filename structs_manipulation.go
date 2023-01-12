package main

import (
    "fmt"
)

type Subscriber struct {
    name string
    amount float64
}

func defaultSubscriber(name string) Subscriber{
    var s Subscriber
    s.name = name
    s.amount = 1.99
    return s
}

func giveDiscount(s Subscriber){
    s.amount = 0.99
}

func giveDiscountV2(s *Subscriber){
    s.amount = 0.99
}

func giveDiscountV3(s Subscriber) Subscriber{
    s.amount = 0.59
    return s
}

func main(){
    var s1 = defaultSubscriber("Saman Silva")
    fmt.Printf("%#v\n", s1)
    giveDiscount(s1)
    fmt.Printf("%#v\n", s1)
    giveDiscountV2(&s1)
    fmt.Printf("%#v\n", s1)
    s1 = giveDiscountV3(s1)
    fmt.Printf("%#v\n", s1)
}