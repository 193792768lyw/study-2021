package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {

	ff := []int(nil)
	arr := append(ff, 33)
	fmt.Println(arr, ff)
	fmt.Println(frequencySort("Aabb"))
}

func frequencySort1(s string) string {
	charMap := map[int32]int{}
	type name struct {
		count int
		char  int32
	}
	arr := make([]*name, 0)
	for _, v := range s {
		if da, ok := charMap[v]; ok {
			arr[da].count += 1
		} else {
			na := &name{
				count: 1,
				char:  v,
			}
			arr = append(arr, na)
			charMap[v] = len(arr) - 1
		}

	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].count > arr[j].count
	})

	bu := strings.Builder{}
	for i := range arr {
		for k := arr[i].count; k > 0; k-- {
			bu.WriteByte(byte(arr[i].char))
		}
	}

	return bu.String()
}

func frequencySort(s string) string {
	cnt := map[byte]int{}
	maxFreq := 0
	for i := range s {
		cnt[s[i]]++
		maxFreq = max(maxFreq, cnt[s[i]])
	}

	buckets := make([][]byte, maxFreq+1)
	for ch, c := range cnt {
		buckets[c] = append(buckets[c], ch)
	}

	ans := make([]byte, 0, len(s))
	for i := maxFreq; i > 0; i-- {
		for _, ch := range buckets[i] {
			ans = append(ans, bytes.Repeat([]byte{ch}, i)...)
		}
	}
	return string(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
