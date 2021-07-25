package main

func main() {

}

func restoreArray(adjacentPairs [][]int) []int {
	res := make([]int, len(adjacentPairs)+1)
	m := map[int][]int{}

	for i := range adjacentPairs {
		v1, v2 := adjacentPairs[i][0], adjacentPairs[i][1]
		m[v1] = append(m[v1], v2)
		m[v2] = append(m[v2], v1)
	}
	for k, v := range m {
		if len(v) == 1 {
			res[0] = k
			res[1] = v[0]
			break
		}
	}
	for i := 2; i < len(res); i++ {
		v := m[res[i-1]]
		if v[0] == res[i-2] {
			res[i] = v[1]
		} else {
			res[i] = v[0]
		}
	}

	return res
}
