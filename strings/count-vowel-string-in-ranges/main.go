package main

import "fmt"

func main() {
	arr := []string{"aba", "bcb", "ece", "aa", "e"}
	queries := [][]int{{0, 2}, {1, 4}, {1, 1}}
	arr2 := []string{"a", "e", "i"}
	queries2 := [][]int{{0, 2}, {0, 1}, {2, 2}}

	fmt.Println(vowelStrings(arr, queries))
	fmt.Println(vowelStrings(arr2, queries2))
}

func vowelStrings(words []string, queries [][]int) []int {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	vowelsCount := make([]int, len(queries))
	prefixSumArr := make([]int, len(words)+1)
	for i := 1; i <= len(words); i++ {
		prefixSumArr[i] = prefixSumArr[i-1]
		if vowels[rune(words[i-1][0])] && vowels[rune(words[i-1][len(words[i-1])-1])] {
			prefixSumArr[i]++
		}
	}
	fmt.Println(prefixSumArr)

	for i, query := range queries {
		vowelsCount[i] = prefixSumArr[query[1]+1] - prefixSumArr[query[0]]
	}

	return vowelsCount
}
