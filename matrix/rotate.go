package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 0, 1, 2},
		{3, 4, 5, 6},
	}
	fmt.Println("Matrix before rotation:", matrix)
	newMatrix := rotateClockWise(matrix)
	fmt.Println("Matrix after +90 degree rotation", newMatrix)
	newMatrix1 := rotateCounterClockWise(matrix)
	fmt.Println("Matrix after -90 degree rotation", newMatrix1)
}

func rotateClockWise(matrix [][]int) [][]int {
	n := len(matrix)
	ret := make([][]int, n, (n*2)+1)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ret[i][j] = matrix[n-j-1][i]
		}
	}
	return ret
}

func rotateCounterClockWise(matrix [][]int) [][]int {
	n := len(matrix)
	ret := make([][]int, n, (n*2)+1)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ret[i][j] = matrix[j][n-i-1]
		}
	}
	return ret
}
