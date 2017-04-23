package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	s := []float64{1, 2, 3, 4, 5, 6, 7}
	var peak float64
	findPeak(s, peak)
}

func findPeak(s float64, peak float64) (float64, float64) {

	for i := 0; i < len(s); i++ {
		if s[i] < s[i+1] {
			peak = s[i+1]
		} else {
			peak = s[i]
		}
	}
	return s / 2, peak
}
