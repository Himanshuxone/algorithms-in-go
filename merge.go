package main

import "fmt"

func main() {
	s := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Println(s)
	fmt.Println(mergesort(s))
}

func mergesort(s []int) []int {
	// create a base return for recursion else the function enters into the error of stackoverflow
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := mergesort(s[:n])
	r := mergesort(s[n:])
	return merge(l, r)
}

func merge(l, r []int) []int {
	ret := make([]int, 0, (len(l)+len(r)+1)*2)
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] >= r[0] {
			ret = append(ret, r[0])
			r = r[1:]
		} else {
			ret = append(ret, l[0])
			l = l[1:]
		}
	}
	return ret
}
