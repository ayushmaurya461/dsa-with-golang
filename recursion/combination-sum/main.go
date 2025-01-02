package main

import "fmt"

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {
	arr := [][]int{}
	combinations(0, &arr, candidates, target, []int{})
	return arr
}

func combinations(ind int, arr *[][]int, candidates []int, target int, comb []int) {
	if ind >= len(candidates) {
		if target == 0 {
			combCopy := make([]int, len(comb))
			copy(combCopy, comb)
			*arr = append(*arr, combCopy)
		}
		return
	}
	if candidates[ind] <= target {
		comb = append(comb, candidates[ind])
		combinations(ind, arr, candidates, target-candidates[ind], comb)
		comb = comb[:len(comb)-1]
	}
	combinations(ind+1, arr, candidates, target, comb)
}
