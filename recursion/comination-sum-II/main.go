package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinations(candidates, target))
}

func combinations(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	arr := [][]int{}
	findCombinations(0, target, candidates, &arr, []int{})
	return arr
}

func findCombinations(ind, target int, candidates []int, ans *[][]int, ds []int) {
	if target == 0 {
		combCopy := make([]int, len(ds))
		copy(combCopy, ds)
		*ans = append(*ans, combCopy)
		return
	}
	for i := ind; i < len(candidates); i++ {
		if i > ind && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] > target {
			break
		}
		ds = append(ds, candidates[i])
		findCombinations(i+1, target-candidates[i], candidates, ans, ds)
		ds = ds[:len(ds)-1]
	}
}
