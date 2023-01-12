package main

import "fmt"

func main() {

	var s1 struct {
		firstname string
		lastname  string
	}

	fmt.Printf("%#v\n", s1)
	s1.firstname = "Kamal"
	s1.lastname = "Perera"
	fmt.Printf("%#v\n", s1)

	var s2 = personBuilder("Amal", "Perera")
	fmt.Printf("%#v\n", s2)

	type person struct {
		firstName string
		lastName  string
		age       int
	}

	p1 := person{
		"Chathurika",
		"Sandarenu",
		38,
	}
	var p2 person
	var p3 person = person{}
	fmt.Printf("Person p1 %#V\n", p1)
	fmt.Printf("Person p2 %#V\n", p2)
	fmt.Printf("Person p3 %#V\n", p3)
}

func personBuilder(fn string, ln string) struct {
	firstName string
	lastName  string
} {
	var p struct {
		firstName string
		lastName  string
	}
	p.firstName = fn
	p.lastName = ln
	return p
}
