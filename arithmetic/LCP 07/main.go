package main

func main() {

}
func numWays1(n int, relation [][]int, k int) (ans int) {
	edges := make([][]int, n)
	for _, r := range relation {
		src, dst := r[0], r[1]
		edges[src] = append(edges[src], dst)
	}
	var dfs func(int, int)
	dfs = func(x, step int) {
		if step == k {
			if x == n-1 {
				ans++
			}
			return
		}
		for _, y := range edges[x] {
			dfs(y, step+1)
		}
	}
	dfs(0, 0)
	return
}

func numWays(n int, relation [][]int, k int) (ans int) {
	edges := make([][]int, n)
	for _, r := range relation {
		src, dst := r[0], r[1]
		edges[src] = append(edges[src], dst)
	}

	step := 0
	q := []int{0}
	for ; len(q) > 0 && step < k; step++ {
		tmp := q
		q = nil
		for _, x := range tmp {
			for _, y := range edges[x] {
				q = append(q, y)
			}
		}
	}

	if step == k {
		for _, x := range q {
			if x == n-1 {
				ans++
			}
		}
	}
	return
}
