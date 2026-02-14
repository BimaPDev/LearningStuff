package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range strs {
		// TODO 1: make counts (26)
		charCount := make([]int, 26)
		for _, char := range word {
			charCount[char-'a']++
		}
		key := fmt.Sprint(charCount)
		groups[key] = append(groups[key], word)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(input))
}
