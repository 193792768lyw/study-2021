package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(displayTable([][]string{
		{"David", "3", "Ceviche"}, {"Corina", "10", "Beef Burrito"}, {"David", "3", "Fried Chicken"},
		{"Carla", "5", "Water"}, {"Carla", "5", "Ceviche"}, {"Rous", "3", "Ceviche"},
	}))
}
func displayTable(orders [][]string) [][]string {
	res := make([][]string, 0)
	res = append(res, []string{"Table"})
	tableMap := make(map[string]struct{})
	tableMap1 := make(map[string]int)
	foodMap1 := make(map[string]int)
	foodMap := make(map[string]struct{})
	foodArr := make([]string, 0)
	tableArr := make([]string, 0)
	for i := range orders {
		if _, ok := tableMap[orders[i][1]]; !ok {
			tableArr = append(tableArr, orders[i][1])
			tableMap[orders[i][1]] = struct{}{}
		}
		if _, ok := foodMap[orders[i][2]]; !ok {
			foodArr = append(foodArr, orders[i][2])
			foodMap[orders[i][2]] = struct{}{}
		}
	}
	sort.Slice(foodArr, func(i, j int) bool {
		return foodArr[i] < foodArr[j]
	})
	sort.Slice(tableArr, func(i, j int) bool {
		a, _ := strconv.Atoi(tableArr[i])
		b, _ := strconv.Atoi(tableArr[j])
		return a < b
	})
	fill := make([]string, 0)
	for i := range foodArr {
		res[0] = append(res[0], foodArr[i])
		foodMap1[foodArr[i]] = len(res[0]) - 1
		fill = append(fill, "0")
	}
	for i := 0; i < len(tableArr); i++ {
		res = append(res, []string{tableArr[i]})
		res[i+1] = append(res[i+1], fill...)
		tableMap1[tableArr[i]] = i + 1

	}
	for _, order := range orders {
		i := tableMap1[order[1]]
		j := foodMap1[order[2]]
		v, _ := strconv.Atoi(res[i][j])
		res[i][j] = strconv.Itoa(v + 1)

	}
	return res
}
