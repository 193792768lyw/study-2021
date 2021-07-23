package main

import "fmt"

func main() {
	fmt.Println(isCovered([][]int{{6, 29}, {33, 43}, {7, 50}, {11, 34}, {24, 26}, {32, 47}}, 1, 50))
}

func isCovered1(ranges [][]int, left int, right int) bool {
	for _, arr := range ranges {
		if left > arr[1] || right < arr[0] {
			continue
		} else {
			if arr[0] <= left && arr[1] >= right {
				return true
			} else if right > arr[1] && left >= arr[0] {
				left = arr[1] + 1
			} else {
				right = arr[0] - 1
			}
		}
	}

	if right < left {
		return true
	}

	return false

}

func isCovered(ranges [][]int, left, right int) bool {
	diff := [52]int{} // 差分数组
	for _, r := range ranges {
		diff[r[0]]++
		diff[r[1]+1]--
	}
	cnt := 0 // 前缀和
	for i := 1; i <= 50; i++ {
		cnt += diff[i]
		if cnt <= 0 && left <= i && i <= right {
			return false
		}
	}
	return true
}
