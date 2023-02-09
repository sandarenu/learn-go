package main

import "fmt"

// Interfaces are usually named with “er” endings
type Animaler interface {
	MakeSound()
}

type Herbivore interface {
	EatGrass()
}

type Carnivore interface {
	EatMeat()
}

// Interface embedding
type Omnivore interface {
	Herbivore
	Carnivore
}

/*
A concrete type does not declare that it implements an interface.
If the method set for a concrete type contains all of the methods
in the method set for an interface, the concrete type implements the interface
*/
type Goat struct {
	Sound string
}

func (g Goat) MakeSound() {
	fmt.Println("Goat says ", g.Sound)
}

func (g Goat) EatGrass() {
	fmt.Println("Goat is eating grass... ")
}

type Human struct {
	Sound string
}

func (g Human) MakeSound() {
	fmt.Println("Human says ", g.Sound)
}

func (g Human) EatGrass() {
	fmt.Println("Human is eating grass... ")
}

func (h Human) EatMeat() {
	fmt.Println("Human is eating meat... ")
}

func main() {

	g := Goat{"bheeee"}
	g.MakeSound()

	var gi Animaler
	gi = g
	gi.MakeSound()

	var hi Herbivore
	hi = g
	hi.EatGrass()

	hm := Human{"Hi"}
	hm.MakeSound()

	var o Omnivore
	o = hm
	o.EatGrass()
	o.EatMeat()

    //Empty Interface
    var animal interface{}
    fmt.Printf("Empty interface %#v\n", animal)
    animal = 1
    fmt.Printf("Empty interface can hold anything %+v\n", animal)
    animal = hi
    animal = hm
    fmt.Printf("Animal %+v\n", animal)
}
