package main

import "fmt"

var countmap = make(map[string]int)

func main() {
	var str string = "aaabbcdadwdwedddaa"

	for i := 0; i < len(str); i++ {
		a := string(str[i])
		if _, ok := countmap[a]; !ok {
			countmap[a] = 1
		} else {
			countmap[a]++
		}

	}
	fmt.Println(countmap)
}
