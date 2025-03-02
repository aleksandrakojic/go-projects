package main

import (
	"fmt"
)

type Animal interface {
	Sound()
}

type Cat struct{}

func (c Cat) Sound()        { fmt.Println("Meow") }
func (c Cat) SpecialToCat() { fmt.Println("Cat special") }

type Dog struct{}

func (d Dog) Sound()       { fmt.Println("Woof") }
func (c Dog) UniqueToDog() { fmt.Println("Dog unique") }

type Domesticated interface {
	Cat | Dog // Not Owls
	Animal
}

// An Owl is an “wild” animal,
// Thus not in the above union of cats and dogs
type Owl struct{}

func (Owl) Sound() { fmt.Println("Owl hoo") }

// Here we limit ourselves to domesticated animals
// If you passed in a 'wild' animal, it would not work
func SoundOff[H Domesticated](animal H) H {
	animal.Sound()
	switch a := any(animal).(type) {
	case Dog:
		a.UniqueToDog()
	case Cat:
		a.SpecialToCat()
	default:
		fmt.Println("Then hoo?")
	}
	return animal
}
func main() {
	var c Cat = SoundOff(Cat{})
	d := SoundOff(Dog{})
	c.Sound()
	c.SpecialToCat()
	d.Sound()
	d.UniqueToDog()
	// SoundOff(Owl{})
}
