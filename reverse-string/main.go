package main

import "fmt"

func main() {
	// create a slice to story binary number for decimal
	var input string
	output := []byte{}
	input = "multitude"
	// use a value to store decimal number
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	fmt.Printf("String:- %s\n", input)
	fmt.Printf("Reverse string:- %s", string(output))
}
