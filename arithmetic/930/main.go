package main

import "fmt"

func main() {
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
}
func numSubarraysWithSum1(nums []int, goal int) int {
	res := 0
	temp := make([]int, 0)
	temp = append(temp, 0)
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			nums[i] += nums[i-1]
		}
		temp = append(temp, nums[i])
	}
	for i := 1; i < len(temp); i++ {
		for k := i; k < len(temp); k++ {
			if temp[k]-temp[i-1] == goal {
				res++
			} else if temp[k]-temp[i-1] > goal {
				break
			}
		}
	}
	return res
}
func numSubarraysWithSum(nums []int, goal int) (ans int) {
	cnt := map[int]int{}
	sum := 0
	for _, num := range nums {
		cnt[sum]++
		sum += num
		ans += cnt[sum-goal]
	}
	return
}
