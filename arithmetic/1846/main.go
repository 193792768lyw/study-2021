package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumElementAfterDecrementingAndRearranging([]int{1, 2, 3, 4, 5}))
}

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	max := 1
	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		if abs(arr[i]-arr[i-1]) > 1 {
			arr[i] = arr[i-1] + 1
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println(arr)
	return max
}

func abs(da int) int {
	if da < 0 {
		return -da
	}
	return da
}
