package main

import "fmt"

func main() {
	s := []int{7, 9, 14, 16, 3, 2, 10, 4, 1, 8}
	fmt.Printf("Unsorted Heap:- %+v \n", s)
	for i := 0; i < len(s); i++ {
		heapify(s, len(s), i)
	}
	heapify(s, len(s), 0)
	fmt.Printf("Sorted Heap:- %+v\n", s)
	s, max := extractMax(s, len(s), 0)
	fmt.Printf("Max element:- %d\n", max)
	fmt.Printf("Array after extraction Heap:- %+v\n", s)
	s, max = extractMax(s, len(s), 0)
	fmt.Printf("Max element:- %d\n", max)
	fmt.Printf("Array after extraction Heap:- %+v\n", s)
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		heapify(arr, n, largest)
	}
}

func extractMax(arr []int, n, i int) ([]int, int) {
	largest := arr[0]
	arr = arr[1:n]
	heapify(arr, n, i)
	return arr, largest
}
