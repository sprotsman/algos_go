package main

import "fmt"

// Leetcode 1. Two Sum

// Given an array of integers nums and an integer target, return indices of the
// two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may
// not use the same element twice.

func main() {
	target := 9
	nums := []int{2, 7, 11, 15}

	fmt.Println(twoSum(nums, target))

	nums = []int{2, 3, 6, 8}
	fmt.Println(twoSum(nums, target))

	target = 6
	nums = []int{3, 3}
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	// create map of value:index
	m := make(map[int]int)

	for i, num := range nums {
		difference := target - num

		// Is the difference (key) in the map?
		d, ok := m[difference]
		if ok {
			// Return the key-index of 'difference' in nums and the index as a
			// slice.
			return []int{d, i}
		}
		// !ok add the value (i) to the map at key (num)
		m[num] = i
	}

	// Return an empty slice if there is no solution.
	return []int{}
}
