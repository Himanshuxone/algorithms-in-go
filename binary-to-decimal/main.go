package main

import "fmt"

func main() {
	// create a slice to story binary number for decimal
	var binary []int
	// use a value to store decimal number
	reverseBinary := binarytodecimal(14)
	for i := len(reverseBinary) - 1; i >= 0; i-- {
		binary = append(binary, reverseBinary[i])
	}
	fmt.Printf("Binary format:- %+v", binary)
}

func binarytodecimal(n int) []int {
	binary := []int{}
	for n >= 1 {
		binary = append(binary, n%2)
		n = n / 2
		if n == 1 {
			binary = append(binary, n)
			break
		}
	}
	return binary
}
