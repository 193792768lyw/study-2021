package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement1(nums []int) int {
	cur := -1
	count := 0

	for i := range nums {
		if i == 0 {
			cur = nums[i]
			count += 1
			continue
		}
		if nums[i] == cur {
			count++
		} else {
			count--
			if count == 0 {
				cur = nums[i]
				count = 1
			}
		}

	}
	t := 0
	for _, num := range nums {
		if num == cur {
			t++
		}
	}
	if t > len(nums)/2 {
		return cur
	}
	return -1
}

func majorityElement(nums []int) int {
	candidate := -1
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}
	if count*2 > len(nums) {
		return candidate
	}
	return -1
}
