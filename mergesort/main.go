package main

import "fmt"

func main() {
	// create an array for merge sort
	arr := []int{7, 9, 14, 16, 3, 2, 10, 4, 1, 8}
	fmt.Printf("Unsorted Array:- %+v\n", arr)
	sortedArray := mergeSort(arr)
	fmt.Printf("Sorted Array:-   %+v", sortedArray)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	arr := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(arr, right...)
		}
		if len(right) == 0 {
			return append(arr, left...)
		}
		if left[0] > right[0] {
			arr = append(arr, right[0])
			right = right[1:]
		} else {
			arr = append(arr, left[0])
			left = left[1:]
		}
	}
	return arr
}
