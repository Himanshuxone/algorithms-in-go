package main

import (
	"fmt"
	"sort"
)

// apply go lang sort interface on custom struc you want to sort
type Person struct {
	Name string
	Age  int
}

// ByAge implement the interface provided by sort package
type ByAge []Person

// Len is the number of elements in the collection.
func (a ByAge) Len() int { return len(a) }

// Less reports whether the element with
// index i should sort before the element with index j.
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Swap swaps the elements with indexes i and j.
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func main() {
	family := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.Sort(ByAge(family))
	fmt.Println(family)
}
