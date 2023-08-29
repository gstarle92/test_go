package main

import "fmt"

//var stringArrayNew = make(map[int]int)

var stringArrayNew = make(map[string]int)

func main() {
	var stringArray = []string{"apple", "mango", "orange", "apple", "apple", "orange", "apple"}
	//var stringArray []int = []int{1, 3, 43, 4, 5, 65, 3, 3, 2, 3, 45, 4, 3, 21, 1}

	for i := 0; i < len(stringArray); i++ {
		fruit := stringArray[i]
		if _, ok := stringArrayNew[fruit]; ok {
			stringArrayNew[fruit]++
		} else {
			stringArrayNew[fruit] = 1
		}

	}
	fmt.Println(stringArrayNew)
}
