package main

import (
	"fmt"
	"sort"
)

// Write a function with an input of list of ints like [1,1,1,3,3,3,3,4,5,6,6,9,9,]
// Return an array that looks like this:
// 3-1, 4-3, 1-4, 1-5, 2-6, 2-9

func main() {
	intSlice := []int{1, 1, 1, 3, 3, 3, 3, 4, 5, 6, 6, 9, 9}

	occurrences(intSlice)
}

func occurrences(s []int) {
	// create a map for element values
	countElement := make(map[int]int)
	for _, element := range s {
		countElement[element] = countElement[element] + 1
	}
	// map[1:3 3:4 4:1 5:1 6:2 9:2]

	// Formatting output of the map

	// get a slice of keys (needed for sort)
	keys := make([]int, 0)

	for key, _ := range countElement {
		keys = append(keys, key)
	}

	// sort the slice
	sort.Ints(keys)  // [1 3 4 5 6 9]

	lastElement := keys[len(countElement)-1]

	for _, key := range keys {
		if key == lastElement {
			fmt.Printf("%d-%d\n", countElement[key], key)
		} else {
			fmt.Printf("%d-%d, ", countElement[key], key)
		}
	}
}