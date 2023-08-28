package main

import "fmt"

func main() {
	var array = [8]int{1, 3, -1, -3, 5, 3, 6, 7}
	var k int = 3
	var max int
	for i := 0; i < len(array); i++ {
		end := i + k
		if end <= len(array) {
			window := array[i:end]
			max = searchMax(window)
			fmt.Println(max)
		}

	}
}
func searchMax(window []int) int {
	max := 0
	for i := 0; i < len(window); i++ {
		if window[i] > max {
			max = window[i]
		}
	}
	return max
}
