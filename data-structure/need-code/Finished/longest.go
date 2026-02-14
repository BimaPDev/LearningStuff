package main

import "fmt"

func longestConsecutive(nums []int) int {
	// first you build the set
	// second you find x-1

	numsHash := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		numsHash[nums[i]] = struct{}{}
	}
	best := 0
	for x := range numsHash {
		if _,ok := numsHash[x-1]; !ok {
			length := 1
			for {
				if _, ok2 := numsHash[x+length]; !ok2 {
					break
				}
				length++
			}
			if length > best {
				best = length
			}
		}
	}
	return best
}

func main() {
	// Basic cases
	fmt.Println(longestConsecutive([]int{}))                 // expected: 0
	fmt.Println(longestConsecutive([]int{1}))                // expected: 1
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2})) // expected: 4 (1,2,3,4)

	// Duplicates + mixed order
	fmt.Println(longestConsecutive([]int{0, 3, 2, 5, 4, 6, 1, 1})) // expected: 7 (0..6)

	// Negative numbers
	fmt.Println(longestConsecutive([]int{-1, -2, -3, 10, 11})) // expected: 3 (-3,-2,-1)
}
