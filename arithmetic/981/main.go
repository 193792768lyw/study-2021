package main

import (
	"fmt"
	"sort"
)

//[[],["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]
func main() {
	obj := Constructor()
	obj.Set("love", "high", 10)
	obj.Set("love", "low", 20)
	fmt.Println(obj.Get("love", 5))
	fmt.Println(obj.Get("love", 10))
	fmt.Println(obj.Get("love", 15))
	fmt.Println(obj.Get("love", 20))
	fmt.Println(obj.Get("love", 25))
}

type pair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	m map[string][]pair
}

func Constructor() TimeMap {
	return TimeMap{map[string][]pair{}}
}

func (m *TimeMap) Set(key string, value string, timestamp int) {
	m.m[key] = append(m.m[key], pair{timestamp, value})
}

func (m *TimeMap) Get(key string, timestamp int) string {
	pairs := m.m[key]
	i := sort.Search(len(pairs), func(i int) bool { return pairs[i].timestamp > timestamp })
	if i > 0 {
		return pairs[i-1].value
	}
	return ""
}
