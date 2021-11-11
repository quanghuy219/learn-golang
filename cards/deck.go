package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Receiver function
// print is a method of type `deck`
// By convention, use one or two letters in the type name as the function variable which is called receiver value
// and avoid use `this` or `self` as the reveiver value
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
