package algorithms

import (
	"fmt"
)

// Check Permutation: Given two strings,write a method to decide if one is a
// permutation of the other.
func CheckPermutation(sample, source string) bool {
	counter := make(map[rune]int)
	for _, c := range sample {
		if _, ok := counter[c]; ok {
			counter[c]++
		} else {
			counter[c] = 1
		}
	}
	for _, c := range source {
		if _, ok := counter[c]; ok {
			counter[c]--
			if counter[c] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	for _, count := range counter {
		if count > 0 {
			return false
		}
	}
	return true
}

func dump(data map[rune]int) {
	for k, v := range data {
		fmt.Printf("%v = %v\n", string(k), v)
	}
}
