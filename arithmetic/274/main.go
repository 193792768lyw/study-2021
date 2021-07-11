package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}

func hIndex(citations []int) int {
	sort.Ints(citations)
	res := 0
	for i := 0; i < len(citations); i++ {
		if citations[i] >= len(citations)-i {
			res = len(citations) - i
			break
		}
	}

	return res
}
