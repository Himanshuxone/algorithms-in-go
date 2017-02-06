package main

import "fmt"

func main() {
	s := []int{38, 27, 43, 3, 9, 82, 10}
	sorted := make([]int, len(s), (len(s)+1)*2)
	sorted = insertionsort(s)
	fmt.Println(sorted)
	//fmt.Println(s[0])
}

func insertionsort(s []int) []int {
	var key, tmp int
	for i := 0; i < len(s); i++ {
		key = s[i]
		for j := i; j > 0; j-- {
			if key < s[j-1] {
				tmp = s[j-1]
				s[j-1] = key
				s[j] = tmp
			}
			fmt.Println(s)
		}
	}
	return s
}
