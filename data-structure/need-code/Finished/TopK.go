package main

import "fmt"

func topKFrequent(nums []int, k int) []int {

	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	res := make([]int, 0, k)
	for i := len(nums); i >= 1; i-- {
		for _, num := range buckets[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	result := topKFrequent(nums, k)
	fmt.Println("Top K Frequent:", result)
}
