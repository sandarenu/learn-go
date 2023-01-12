package main

import (
	"fmt"
)

type SubscriberUser struct {
	name    string
	amount  float64
	address Address
}

type Address struct {
	homeNo string
	street string
	city   string
}

//Struct Embedding
type SubscriberUserV2 struct {
	name   string
	amount float64
	Address
	city string
}

func main() {
	s1 := SubscriberUser{name: "Shane Perera", amount: 2.99, address: Address{homeNo: "123/5/A", street: "Temple Road", city: "Kadawatha"}}
	fmt.Printf("%#v\n", s1)

	s2 := SubscriberUser{}
	fmt.Printf("%#v\n", s2)
	s2.name = "Ruwan Perera"
	s2.address.homeNo = "34/5"
	s2.address.street = "Church Road"
	s2.address.city = "Ekala"
	fmt.Printf("%#v\n", s2)

	s3 := SubscriberUserV2{}
	fmt.Printf("%#v\n", s3)
	s3.name = "Harry Perera"
	s3.homeNo = "13/4"
	s3.street = "Temple Road"
	s3.city = "Kegalla1"
	s3.Address.city = "Kegalla"
	fmt.Printf("%#v\n", s3)
}
