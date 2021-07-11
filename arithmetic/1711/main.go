package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(countPairs([]int{1, 1, 1, 3, 3, 3, 7}))
}
func countPairs1(deliciousness []int) int {
	sort.Ints(deliciousness)
	tar := map[int]struct{}{1: {}}
	cur := 1
	index := 1
	vv := 0
	if len(deliciousness) == 1 {
		vv = deliciousness[len(deliciousness)-1]
	} else {
		vv = deliciousness[len(deliciousness)-1] + deliciousness[len(deliciousness)-2]

	}
	for cur <= vv {
		cur = int(math.Pow(float64(2), float64(index)))
		tar[cur] = struct{}{}
		index++
	}

	res := 0
	for i := 0; i < len(deliciousness); i++ {
		for k := i + 1; k < len(deliciousness); k++ {
			if _, ok := tar[deliciousness[i]+deliciousness[k]]; ok {
				res++
			}
		}
	}
	return res % (1e9 + 7)
}

func countPairs(deliciousness []int) (ans int) {
	maxVal := deliciousness[0]
	for _, val := range deliciousness[1:] {
		if val > maxVal {
			maxVal = val
		}
	}
	maxSum := maxVal * 2
	cnt := map[int]int{}
	for _, val := range deliciousness {
		for sum := 1; sum <= maxSum; sum <<= 1 {
			ans += cnt[sum-val]
		}
		cnt[val]++
	}
	return ans % (1e9 + 7)
}
