package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(maxFrequency([]int{3, 9, 6}, 2))
}

func maxFrequency1(nums []int, k int) int {

	res := math.MinInt64
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		te := nums[i]
		tk := k
		rt := 1
		index := i - 1
		for tk > 0 && index >= 0 {
			if tk-te+nums[index] >= 0 {
				rt++
			} else {
				break
			}

			tk = tk - te + nums[index]
			index--
		}
		if res < rt {
			res = rt
		}
	}
	return res
}

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	ans := 1
	for l, r, total := 0, 1, 0; r < len(nums); r++ {
		total += (nums[r] - nums[r-1]) * (r - l)
		for total > k {
			total -= nums[r] - nums[l]
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
