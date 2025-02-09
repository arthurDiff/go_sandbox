package main

import (
	"fmt"
)

func test(_ []int, _ int) int {
	return -1
}

func max(nums []int) (int, int) {
	idx, max := 0, nums[0]
	for i_idx, v := range nums {
		if max > v {
			continue
		}
		idx = i_idx
		max = v
	}
	return idx, max
}

func main() {
	fmt.Println(test([]int{1, 2, 3}, 1))
}
