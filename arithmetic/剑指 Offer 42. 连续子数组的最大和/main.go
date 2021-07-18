package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-1, 0, -2}))
}

func maxSubArray(nums []int) int {
	res := math.MinInt64

	cur := 0
	for _, v := range nums {
		if v+cur < v {
			cur = v
		} else {
			cur += v
		}
		if res < cur {
			res = cur
		}
	}

	return res
}
