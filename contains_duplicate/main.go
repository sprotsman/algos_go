package main

import "fmt"

// Leetcode 217. Contains Duplicate

func main() {
	nums := []int{1,1,1,3,3,4,3,2,4,2}
	fmt.Printf("%d contains duplicates? → ", nums)
	fmt.Println(containsDuplicate(nums))

	nums = []int{1,2,3,4}
	fmt.Printf("%d contains duplicates? → ", nums)
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	set := make(map[int]int)

	for _, v := range nums {
		set[v] = set[v] + 1
	}

	return len(nums) != len(set)
}
