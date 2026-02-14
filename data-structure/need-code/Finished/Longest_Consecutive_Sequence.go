package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	seen := make(map[int]bool, len(nums))
	for _, n := range nums {
		seen[n] = true
	}

	maxLen := 0

	for n := range seen {
		// TODO:
		// 1) Check if n is the start of a sequence (n-1 not in map)
		// 2) Count upward while consecutive numbers exist
		// 3) Update maxLen with the best streak length
		_ = n
	}

	return maxLen
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println("Input:", nums)
	fmt.Println("Longest consecutive length:", longestConsecutive(nums))
}
