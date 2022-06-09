package main

import (
	"fmt"
	"sort"
	"strings"
)

/* 242. Valid Anagram

Given two strings s and t, return true if t is an anagram of s, and false
otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different
word or phrase, typically using all the original letters exactly once.
*/

func main() {
	// Example 1:
	s := "anagram"
	t := "nagaram"
	// Output: true

	if len(s) != len(t) {
		return
	}

	fmt.Println(isAnagram(s, t))
	fmt.Println(isAnagram2(s, t))

	// Example 2:
	s = "rat"
	t = "car"
	// Output: false
	fmt.Println(isAnagram(s, t))
	fmt.Println(isAnagram2(s, t))
}

func isAnagram(s string, t string) bool {
	if sortString(s) == sortString(t) {
		return true
	}
	return false
}

func isAnagram2(s string, t string) bool {
	return strings.Contains(sortString(s), sortString(t))
}

func sortString(s string) string {
	str := strings.Split(s, "")
	sort.Strings(str)
	return strings.Join(str, "")
}
