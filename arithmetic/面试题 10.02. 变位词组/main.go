package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams1(strs []string) [][]string {
	s := map[string][]string{}
	for _, str := range strs {
		s1 := []byte(str)
		sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
		sortedStr := string(s1)

		s[sortedStr] = append(s[sortedStr], str)
	}
	res := make([][]string, 0, len(s))
	for _, v := range s {
		res = append(res, v)
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
