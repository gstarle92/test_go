package main

import (
	"fmt"
	"sort"
)

func main() {
	var str1 string = "anagram"
	var str2 string = "naramga"
	if len(str1) == len(str2) {
		str1rune := []rune(str1)
		str2rune := []rune(str2)
		sort.Slice(str1rune, func(i, j int) bool {
			return str1rune[i] < str1rune[j]
		})

		sort.Slice(str2rune, func(i, j int) bool {
			return str2rune[i] < str2rune[j]
		})
		if string(str1rune) == string(str2rune) {
			fmt.Println("this are perfect anagram string")
		}
		// fmt.Println("Sorted runes:", string(str1rune))

		// fmt.Println("Sorted runes:", string(str2rune))

	}
}
