package main

import "fmt"

func main() {
	// create an adjacency list or maps of strings with values holding the connected nodes
	// in case
	maps := map[string]map[int]string{
		"a": {1: "z", 2: "s"},
		"z": {1: "a"},
		"s": {1: "a", 2: "x"},
		"x": {1: "d", 2: "s", 3: "c"},
		"d": {1: "c", 2: "x", 3: "f"},
		"c": {1: "x", 2: "d", 3: "f", 4: "v"},
		"v": {1: "c", 2: "f"},
		"f": {1: "d", 2: "c", 3: "v"},
	}

	fmt.Printf("Adjacency List:- %+v", maps)
}
