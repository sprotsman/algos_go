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
	fmt.Println(isAnagram(s, t))
	fmt.Println(isAnagram2(s, t))

	// Example 2:
	s = "rat"
	t = "car"
	// Output: false
	fmt.Println(isAnagram(s, t))
	fmt.Println(isAnagram2(s, t))

	// Example 3:
	s = "stop"
	t = "pots"
	fmt.Println(isAnagram3(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if sortString(s) == sortString(t) {
		return true
	}
	return false
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	return strings.Contains(sortString(s), sortString(t))
}

func isAnagram3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	// fmt.Println("m:", m)

	// if value doesn't exist, add (increment (1)) to key in map
	for _, v := range s {
		m[v]++
	}

	// if value does exist, decrement the key's value in map
	for _, v := range t {
		m[v]--
	}

	// if key's values in map are not all 0, return false
	for k, _ := range m {
		if m[k] != 0 {
			return false
		}
	}

	return true
}

func sortString(s string) string {
	str := strings.Split(s, "")
	sort.Strings(str)
	return strings.Join(str, "")
}
