package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minPairSum([]int{3, 5, 4, 2, 4, 6}))
}

func minPairSum(nums []int) int {
	sort.Ints(nums)
	res := math.MinInt64

	left := 0
	right := len(nums) - 1
	for left < right {
		s := nums[left] + nums[right]
		if s > res {
			res = s
		}
		left++
		right--
	}
	return res
}
