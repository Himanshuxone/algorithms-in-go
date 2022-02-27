package main

import "fmt"

func main() {
	s := []int{7, 9, 14, 16, 3, 2, 10, 4, 1, 8}
	length := len(s)/2 - 1
	var sorted = make([]int, length, length*2+1)
	for i := length; i >= 0; i-- {
		sorted = heapify(s, i)
		fmt.Println(sorted)
	}
}

func heapify(s []int, root int) []int {
	var largest, tmp, left, right int

	left = 2*root + 1
	right = 2*root + 2

	if s[root] > s[left] && s[root] > s[right] {
		return s
	}

	if s[root] < s[left] {
		largest = left
	} else if right < len(s) && s[root] < s[right] && s[left] < s[right] {
		largest = right
	}

	//if largest != 0 {
	tmp = s[largest]
	s[largest] = s[root]
	s[root] = tmp
	//fmt.Printf("heap is %d \n", s)
	if largest <= len(s)/2-1 {
		heapify(s, largest)
	}
	return s
	//if root == 0 {
	//	return s
	//}
	//}
}
